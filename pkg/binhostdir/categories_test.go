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
package binhostdir_test

import (
	"regexp"

	. "github.com/Sabayon/pkgs-checker/pkg/binhostdir"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binhost Dir", func() {

	Context("Check Categories", func() {

		validCats := []string{
			"acct-group",
			"acct-user",
			"app-accessibility",
			"app-admin",
			"app-antivirus",
			"app-arch",
			"app-backup",
			"app-benchmarks",
			"app-cdr",
			"app-crypt",
			"app-dicts",
			"app-doc",
			"app-editors",
			"app-emacs",
			"app-emulation",
			"app-eselect",
			"app-forensics",
			"app-i18n",
			"app-laptop",
			"app-leechcraft",
			"app-metrics",
			"app-misc",
			"app-mobilephone",
			"app-office",
			"app-officeext",
			"app-pda",
			"app-portage",
			"app-shells",
			"app-text",
			"app-vim",
			"app-xemacs",
			"dev-ada",
			"dev-cpp",
			"dev-db",
			"dev-dotnet",
			"dev-embedded",
			"dev-erlang",
			"dev-games",
			"dev-go",
			"dev-haskell",
			"dev-java",
			"dev-lang",
			"dev-libs",
			"dev-lisp",
			"dev-lua",
			"dev-ml",
			"dev-perl",
			"dev-php",
			"dev-python",
			"dev-qt",
			"dev-ros",
			"dev-ruby",
			"dev-scheme",
			"dev-tcltk",
			"dev-tex",
			"dev-texlive",
			"dev-util",
			"dev-vcs",
			"games-action",
			"games-arcade",
			"games-board",
			"games-emulation",
			"games-engines",
			"games-fps",
			"games-kids",
			"games-misc",
			"games-mud",
			"games-puzzle",
			"games-roguelike",
			"games-rpg",
			"games-server",
			"games-simulation",
			"games-sports",
			"games-strategy",
			"games-util",
			"gnome-base",
			"gnome-extra",
			"gnustep-apps",
			"gnustep-base",
			"gnustep-libs",
			"gui-apps",
			"gui-libs",
			"gui-wm",
			"java-virtuals",
			"kde-apps",
			"kde-frameworks",
			"kde-misc",
			"kde-plasma",
			"lxde-base",
			"lxqt-base",
			"mail-client",
			"mail-filter",
			"mail-mta",
			"mate-base",
			"mate-extra",
			"media-fonts",
			"media-gfx",
			"media-libs",
			"media-plugins",
			"media-radio",
			"media-sound",
			"media-tv",
			"media-video",
			"net-analyzer",
			"net-dialup",
			"net-dns",
			"net-firewall",
			"net-fs",
			"net-ftp",
			"net-im",
			"net-irc",
			"net-libs",
			"net-mail",
			"net-misc",
			"net-nds",
			"net-news",
			"net-nntp",
			"net-p2p",
			"net-print",
			"net-proxy",
			"net-voip",
			"net-vpn",
			"net-wireless",
			"perl-core",
			"ros-meta",
			"sci-astronomy",
			"sci-biology",
			"sci-calculators",
			"sci-chemistry",
			"sci-electronics",
			"sci-geosciences",
			"sci-libs",
			"sci-mathematics",
			"sci-misc",
			"sci-physics",
			"sci-visualization",
			"sec-policy",
			"sys-apps",
			"sys-auth",
			"sys-block",
			"sys-boot",
			"sys-cluster",
			"sys-devel",
			"sys-fabric",
			"sys-firmware",
			"sys-fs",
			"sys-kernel",
			"sys-libs",
			"sys-power",
			"sys-process",
			"virtual",
			"www-apache",
			"www-apps",
			"www-client",
			"www-misc",
			"www-plugins",
			"www-servers",
			"x11-apps",
			"x11-base",
			"x11-drivers",
			"x11-libs",
			"x11-misc",
			"x11-plugins",
			"x11-terms",
			"x11-themes",
			"x11-wm",
			"xfce-base",
			"xfce-extra",
		}

		invalidCats := []string{
			"skel.ebuild",
			"skel.metadata.xml",
			"distfiles",
			"eclass",
			"scripts",
			"header.txt",
			"Manifest",
			"Manifest.files.gz",
			"licenses",
			"local",
			"metadata",
			"packages",
			"profiles",
		}

		var regexCat = regexp.MustCompile(RegexCatString)

		for _, cat := range validCats {

			match := regexCat.MatchString(cat)
			It("Check cat "+cat, func() {
				Expect(match).Should(Equal(true))
			})
		}

		for _, cat := range invalidCats {

			match := regexCat.MatchString(cat)
			It("Check cat "+cat, func() {
				Expect(match).Should(Equal(false))
			})
		}
	})
})