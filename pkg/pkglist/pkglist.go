/*

Copyright (C) 2017-2019  Daniele Rondina <geaaru@sabayonlinux.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/
package pkglist

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	logger "github.com/sirupsen/logrus"

	"github.com/Sabayon/pkgs-checker/pkg/binhostdir"
)

func PkgListParser(data []byte) ([]string, error) {
	var ans []string = make([]string, 0)

	reader := bytes.NewBuffer(data)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return ans, err
		}
		ans = append(ans, strings.TrimRight(line, "\r\n"))
	}

	return ans, nil
}

func PkgListCreate(binhostDir string, log *logger.Logger) ([]string, error) {
	// TODO: handle logger outside
	ans := make([]string, 0)

	if binhostDir == "" {
		return ans, errors.New("Invalid binhostDir")
	}

	binHostTree := make(map[string][]string, 0)

	err := binhostdir.AnalyzeBinHostDirectory(binhostDir, log, &binHostTree)
	if err != nil {
		return ans, err
	}

	if len(binHostTree) > 0 {
		keys := make([]string, 0)
		for k, _ := range binHostTree {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, cat := range keys {
			pkgs := binHostTree[cat]
			sort.Strings(pkgs)

			for _, p := range pkgs {
				f := filepath.Base(p)
				ans = append(ans,
					fmt.Sprintf("%s/%s",
						cat, f[0:strings.Index(f, filepath.Ext(f))]))
			}
		}
	}

	return ans, nil
}

func PkgListWriteFile(pkgs []string, f string) error {
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	err = PkgListWrite(pkgs, w)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}

func PkgListWrite(pkgs []string, out io.Writer) error {
	for _, p := range pkgs {
		_, err := io.WriteString(out, fmt.Sprintf("%s\n", p))
		if err != nil {
			return err
		}
	}

	return nil
}