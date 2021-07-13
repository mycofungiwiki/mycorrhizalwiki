// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "path/filepath"

//line views/stuff.qtpl:2
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/stuff.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/stuff.qtpl:5
import "github.com/bouncepaw/mycorrhiza/util"

//line views/stuff.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:7
func StreamBaseHTML(qw422016 *qt422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:7
	qw422016.N().S(`
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>`)
//line views/stuff.qtpl:13
	qw422016.E().S(title)
//line views/stuff.qtpl:13
	qw422016.N().S(`</title>
		<link rel="shortcut icon" href="/static/favicon.ico">
		<link rel="stylesheet" href="/static/style.css">
		<script src="/static/shortcuts.js"></script>
		`)
//line views/stuff.qtpl:17
	for _, el := range headElements {
//line views/stuff.qtpl:17
		qw422016.N().S(el)
//line views/stuff.qtpl:17
	}
//line views/stuff.qtpl:17
	qw422016.N().S(`
	</head>
	<body>
		<header>
			<nav class="main-width top-bar">
				<ul class="top-bar__wrapper">
					<li class="top-bar__section top-bar__section_home">
						<a class="top-bar__home-link" href="/">`)
//line views/stuff.qtpl:24
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:24
	qw422016.N().S(`</a>
					</li>
					<li class="top-bar__section top-bar__section_search">
						<form class="top-bar__search" method="GET" action="/title-search">
							<input type="text" name="q" placeholder="Title search" class="top-bar__search-bar">
						</form>
					</li>
					`)
//line views/stuff.qtpl:31
	if cfg.UseAuth {
//line views/stuff.qtpl:31
		qw422016.N().S(`
					<li class="top-bar__section top-bar__section_auth">
						<ul class="top-bar__auth auth-links">
							<li class="auth-links__box auth-links__user-box">
								`)
//line views/stuff.qtpl:35
		if u.Group == "anon" {
//line views/stuff.qtpl:35
			qw422016.N().S(`
								<a href="/login" class="auth-links__login-link">Login</a>
								`)
//line views/stuff.qtpl:37
		} else {
//line views/stuff.qtpl:37
			qw422016.N().S(`
								<a href="/hypha/`)
//line views/stuff.qtpl:38
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:38
			qw422016.N().S(`/`)
//line views/stuff.qtpl:38
			qw422016.E().S(u.Name)
//line views/stuff.qtpl:38
			qw422016.N().S(`" class="auth-links__user-link">`)
//line views/stuff.qtpl:38
			qw422016.E().S(util.BeautifulName(u.Name))
//line views/stuff.qtpl:38
			qw422016.N().S(`</a>
								`)
//line views/stuff.qtpl:39
		}
//line views/stuff.qtpl:39
		qw422016.N().S(`
							</li>
							`)
//line views/stuff.qtpl:41
		if cfg.AllowRegistration && u.Group == "anon" {
//line views/stuff.qtpl:41
			qw422016.N().S(`
							<li class="auth-links__box auth-links__register-box">
								<a href="/register" class="auth-links__register-link">Register</a>
							</li>
							`)
//line views/stuff.qtpl:45
		}
//line views/stuff.qtpl:45
		qw422016.N().S(`
							`)
//line views/stuff.qtpl:46
		if u.Group == "admin" {
//line views/stuff.qtpl:46
			qw422016.N().S(`
							<li class="auth-links__box auth-links__admin-box">
								<a href="/admin" class="auth-links__admin-link">Admin panel</a>
							</li>
							`)
//line views/stuff.qtpl:50
		}
//line views/stuff.qtpl:50
		qw422016.N().S(`
						</ul>
					</li>
					`)
//line views/stuff.qtpl:53
	}
//line views/stuff.qtpl:53
	qw422016.N().S(`
					<li class="top-bar__section top-bar__section_highlights">
						<ul class="top-bar__highlights">
`)
//line views/stuff.qtpl:56
	for _, link := range cfg.HeaderLinks {
//line views/stuff.qtpl:56
		qw422016.N().S(`						`)
//line views/stuff.qtpl:57
		if link.Href != "/" {
//line views/stuff.qtpl:57
			qw422016.N().S(`
							<li class="top-bar__highlight">
								<a class="top-bar__highlight-link" href="`)
//line views/stuff.qtpl:59
			qw422016.E().S(link.Href)
//line views/stuff.qtpl:59
			qw422016.N().S(`">`)
//line views/stuff.qtpl:59
			qw422016.E().S(link.Display)
//line views/stuff.qtpl:59
			qw422016.N().S(`</a>
							</li>
						`)
//line views/stuff.qtpl:61
		}
//line views/stuff.qtpl:61
		qw422016.N().S(`
`)
//line views/stuff.qtpl:62
	}
//line views/stuff.qtpl:62
	qw422016.N().S(`						</ul>
					</li>
				</ul>
			</nav>
		</header>
		`)
//line views/stuff.qtpl:68
	qw422016.N().S(body)
//line views/stuff.qtpl:68
	qw422016.N().S(`
		<template id="dialog-template">
			<div class="dialog-backdrop"></div>
			<div class="dialog" tabindex="0">
				<div class="dialog__header">
					<h1 class="dialog__title"></h1>
					<button class="dialog__close-button" aria-label="Close this dialog"></button>
				</div>

				<div class="dialog__content"></div>
			</div>
		</template>
		`)
//line views/stuff.qtpl:80
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:80
	qw422016.N().S(`
	</body>
</html>
`)
//line views/stuff.qtpl:83
}

//line views/stuff.qtpl:83
func WriteBaseHTML(qq422016 qtio422016.Writer, title, body string, u *user.User, headElements ...string) {
//line views/stuff.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:83
	StreamBaseHTML(qw422016, title, body, u, headElements...)
//line views/stuff.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:83
}

//line views/stuff.qtpl:83
func BaseHTML(title, body string, u *user.User, headElements ...string) string {
//line views/stuff.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:83
	WriteBaseHTML(qb422016, title, body, u, headElements...)
//line views/stuff.qtpl:83
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:83
	return qs422016
//line views/stuff.qtpl:83
}

//line views/stuff.qtpl:85
func StreamTitleSearchHTML(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:85
	qw422016.N().S(`
<div class="layout">
<main class="main-width title-search">
	<h1>Search results for ‘`)
//line views/stuff.qtpl:88
	qw422016.E().S(query)
//line views/stuff.qtpl:88
	qw422016.N().S(`’</h1>
	<p>Every hypha name has been compared with the query. Hyphae that have matched the query are listed below.</p>
	<ul class="title-search__results">
	`)
//line views/stuff.qtpl:91
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:91
		qw422016.N().S(`
		<li class="title-search__entry">
			<a class="title-search__link wikilink" href="/hypha/`)
//line views/stuff.qtpl:93
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:93
		qw422016.N().S(`">`)
//line views/stuff.qtpl:93
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:93
		qw422016.N().S(`</a>
		</li>
	`)
//line views/stuff.qtpl:95
	}
//line views/stuff.qtpl:95
	qw422016.N().S(`
</main>
</div>
`)
//line views/stuff.qtpl:98
}

//line views/stuff.qtpl:98
func WriteTitleSearchHTML(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:98
	StreamTitleSearchHTML(qw422016, query, generator)
//line views/stuff.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:98
}

//line views/stuff.qtpl:98
func TitleSearchHTML(query string, generator func(string) <-chan string) string {
//line views/stuff.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:98
	WriteTitleSearchHTML(qb422016, query, generator)
//line views/stuff.qtpl:98
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:98
	return qs422016
//line views/stuff.qtpl:98
}

// It outputs a poorly formatted JSON, but it works and is valid.

//line views/stuff.qtpl:101
func StreamTitleSearchJSON(qw422016 *qt422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:101
	qw422016.N().S(`
`)
//line views/stuff.qtpl:103
	// Lol
	counter := 0

//line views/stuff.qtpl:105
	qw422016.N().S(`
{
	"source_query": "`)
//line views/stuff.qtpl:107
	qw422016.E().S(query)
//line views/stuff.qtpl:107
	qw422016.N().S(`",
	"results": [
	`)
//line views/stuff.qtpl:109
	for hyphaName := range generator(query) {
//line views/stuff.qtpl:109
		qw422016.N().S(`
		`)
//line views/stuff.qtpl:110
		if counter > 0 {
//line views/stuff.qtpl:110
			qw422016.N().S(`, `)
//line views/stuff.qtpl:110
		}
//line views/stuff.qtpl:110
		qw422016.N().S(`{
			"canonical_name": "`)
//line views/stuff.qtpl:111
		qw422016.E().S(hyphaName)
//line views/stuff.qtpl:111
		qw422016.N().S(`",
			"beautiful_name": "`)
//line views/stuff.qtpl:112
		qw422016.E().S(util.BeautifulName(hyphaName))
//line views/stuff.qtpl:112
		qw422016.N().S(`",
			"url": "`)
//line views/stuff.qtpl:113
		qw422016.E().S(cfg.URL + "/hypha/" + hyphaName)
//line views/stuff.qtpl:113
		qw422016.N().S(`"
		}`)
//line views/stuff.qtpl:114
		counter++

//line views/stuff.qtpl:114
		qw422016.N().S(`
	`)
//line views/stuff.qtpl:115
	}
//line views/stuff.qtpl:115
	qw422016.N().S(`
	]
}
`)
//line views/stuff.qtpl:118
}

//line views/stuff.qtpl:118
func WriteTitleSearchJSON(qq422016 qtio422016.Writer, query string, generator func(string) <-chan string) {
//line views/stuff.qtpl:118
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:118
	StreamTitleSearchJSON(qw422016, query, generator)
//line views/stuff.qtpl:118
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:118
}

//line views/stuff.qtpl:118
func TitleSearchJSON(query string, generator func(string) <-chan string) string {
//line views/stuff.qtpl:118
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:118
	WriteTitleSearchJSON(qb422016, query, generator)
//line views/stuff.qtpl:118
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:118
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:118
	return qs422016
//line views/stuff.qtpl:118
}

//line views/stuff.qtpl:120
func StreamHelpHTML(qw422016 *qt422016.Writer, content string) {
//line views/stuff.qtpl:120
	qw422016.N().S(`
<div class="layout">
<main class="main-width help">
	<article>
	`)
//line views/stuff.qtpl:124
	qw422016.N().S(content)
//line views/stuff.qtpl:124
	qw422016.N().S(`
	</article>
</main>
`)
//line views/stuff.qtpl:127
	qw422016.N().S(helpTopicsHTML())
//line views/stuff.qtpl:127
	qw422016.N().S(`
</div>
`)
//line views/stuff.qtpl:129
}

//line views/stuff.qtpl:129
func WriteHelpHTML(qq422016 qtio422016.Writer, content string) {
//line views/stuff.qtpl:129
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:129
	StreamHelpHTML(qw422016, content)
//line views/stuff.qtpl:129
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:129
}

//line views/stuff.qtpl:129
func HelpHTML(content string) string {
//line views/stuff.qtpl:129
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:129
	WriteHelpHTML(qb422016, content)
//line views/stuff.qtpl:129
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:129
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:129
	return qs422016
//line views/stuff.qtpl:129
}

//line views/stuff.qtpl:131
func StreamHelpEmptyErrorHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:131
	qw422016.N().S(`
<h1>This entry does not exist!</h1>
<p>Try finding a different entry that would help you.</p>
<p>If you want to write this entry by yourself, consider <a class="wikilink wikilink_external wikilink_https" href="https://github.com/bouncepaw/mycorrhiza">contributing</a> to Mycorrhiza Wiki directly.</p>
`)
//line views/stuff.qtpl:135
}

//line views/stuff.qtpl:135
func WriteHelpEmptyErrorHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:135
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:135
	StreamHelpEmptyErrorHTML(qw422016)
//line views/stuff.qtpl:135
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:135
}

//line views/stuff.qtpl:135
func HelpEmptyErrorHTML() string {
//line views/stuff.qtpl:135
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:135
	WriteHelpEmptyErrorHTML(qb422016)
//line views/stuff.qtpl:135
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:135
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:135
	return qs422016
//line views/stuff.qtpl:135
}

//line views/stuff.qtpl:137
func streamhelpTopicsHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:137
	qw422016.N().S(`
<aside class="help-topics layout-card">
	<h2 class="layout-card__title">Help topics</h2>
	<ul class="help-topics__list">
		<li>
			<a href="/help/en">Main</a>
		</li>
		<li>
			<a href="/help/en/hypha">Hypha</a>
			<ul>
				<li>
					<a href="/help/en/attachment">Attachment</a>
				</li>
			</ul
		</li>
	</ul>
</aside>
`)
//line views/stuff.qtpl:154
}

//line views/stuff.qtpl:154
func writehelpTopicsHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:154
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:154
	streamhelpTopicsHTML(qw422016)
//line views/stuff.qtpl:154
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:154
}

//line views/stuff.qtpl:154
func helpTopicsHTML() string {
//line views/stuff.qtpl:154
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:154
	writehelpTopicsHTML(qb422016)
//line views/stuff.qtpl:154
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:154
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:154
	return qs422016
//line views/stuff.qtpl:154
}

//line views/stuff.qtpl:156
func StreamUserListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:156
	qw422016.N().S(`
<div class="layout">
<main class="main-width user-list">
	<h1>List of users</h1>
`)
//line views/stuff.qtpl:161
	var (
		admins     = make([]string, 0)
		moderators = make([]string, 0)
		editors    = make([]string, 0)
	)
	for u := range user.YieldUsers() {
		switch u.Group {
		case "admin":
			admins = append(admins, u.Name)
		case "moderator":
			moderators = append(moderators, u.Name)
		case "editor", "trusted":
			editors = append(editors, u.Name)
		}
	}

//line views/stuff.qtpl:176
	qw422016.N().S(`
	<section>
		<h2>Admins</h2>
		<ol>`)
//line views/stuff.qtpl:179
	for _, name := range admins {
//line views/stuff.qtpl:179
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:180
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:180
		qw422016.N().S(`/`)
//line views/stuff.qtpl:180
		qw422016.E().S(name)
//line views/stuff.qtpl:180
		qw422016.N().S(`">`)
//line views/stuff.qtpl:180
		qw422016.E().S(name)
//line views/stuff.qtpl:180
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:181
	}
//line views/stuff.qtpl:181
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Moderators</h2>
		<ol>`)
//line views/stuff.qtpl:185
	for _, name := range moderators {
//line views/stuff.qtpl:185
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:186
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:186
		qw422016.N().S(`/`)
//line views/stuff.qtpl:186
		qw422016.E().S(name)
//line views/stuff.qtpl:186
		qw422016.N().S(`">`)
//line views/stuff.qtpl:186
		qw422016.E().S(name)
//line views/stuff.qtpl:186
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:187
	}
//line views/stuff.qtpl:187
	qw422016.N().S(`</ol>
	</section>
	<section>
		<h2>Editors</h2>
		<ol>`)
//line views/stuff.qtpl:191
	for _, name := range editors {
//line views/stuff.qtpl:191
		qw422016.N().S(`
			<li><a href="/hypha/`)
//line views/stuff.qtpl:192
		qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:192
		qw422016.N().S(`/`)
//line views/stuff.qtpl:192
		qw422016.E().S(name)
//line views/stuff.qtpl:192
		qw422016.N().S(`">`)
//line views/stuff.qtpl:192
		qw422016.E().S(name)
//line views/stuff.qtpl:192
		qw422016.N().S(`</a></li>
		`)
//line views/stuff.qtpl:193
	}
//line views/stuff.qtpl:193
	qw422016.N().S(`</ol>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:197
}

//line views/stuff.qtpl:197
func WriteUserListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:197
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:197
	StreamUserListHTML(qw422016)
//line views/stuff.qtpl:197
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:197
}

//line views/stuff.qtpl:197
func UserListHTML() string {
//line views/stuff.qtpl:197
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:197
	WriteUserListHTML(qb422016)
//line views/stuff.qtpl:197
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:197
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:197
	return qs422016
//line views/stuff.qtpl:197
}

//line views/stuff.qtpl:199
func StreamHyphaListHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:199
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>List of hyphae</h1>
	<p>This wiki has `)
//line views/stuff.qtpl:203
	qw422016.N().D(hyphae.Count())
//line views/stuff.qtpl:203
	qw422016.N().S(` hyphae.</p>
	<ul class="hypha-list">
		`)
//line views/stuff.qtpl:205
	for h := range hyphae.YieldExistingHyphae() {
//line views/stuff.qtpl:205
		qw422016.N().S(`
		<li class="hypha-list__entry">
			<a class="hypha-list__link" href="/hypha/`)
//line views/stuff.qtpl:207
		qw422016.E().S(h.Name)
//line views/stuff.qtpl:207
		qw422016.N().S(`">`)
//line views/stuff.qtpl:207
		qw422016.E().S(util.BeautifulName(h.Name))
//line views/stuff.qtpl:207
		qw422016.N().S(`</a>
			`)
//line views/stuff.qtpl:208
		if h.BinaryPath != "" {
//line views/stuff.qtpl:208
			qw422016.N().S(`
			<span class="hypha-list__amnt-type">`)
//line views/stuff.qtpl:209
			qw422016.E().S(filepath.Ext(h.BinaryPath)[1:])
//line views/stuff.qtpl:209
			qw422016.N().S(`</span>
			`)
//line views/stuff.qtpl:210
		}
//line views/stuff.qtpl:210
		qw422016.N().S(`
		</li>
		`)
//line views/stuff.qtpl:212
	}
//line views/stuff.qtpl:212
	qw422016.N().S(`
	</ul>
</main>
</div>
`)
//line views/stuff.qtpl:216
}

//line views/stuff.qtpl:216
func WriteHyphaListHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:216
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:216
	StreamHyphaListHTML(qw422016)
//line views/stuff.qtpl:216
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:216
}

//line views/stuff.qtpl:216
func HyphaListHTML() string {
//line views/stuff.qtpl:216
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:216
	WriteHyphaListHTML(qb422016)
//line views/stuff.qtpl:216
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:216
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:216
	return qs422016
//line views/stuff.qtpl:216
}

//line views/stuff.qtpl:218
func StreamAboutHTML(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:218
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<section>
		<h1>About `)
//line views/stuff.qtpl:222
	qw422016.E().S(cfg.WikiName)
//line views/stuff.qtpl:222
	qw422016.N().S(`</h1>
		<ul>
			<li><b><a href="https://mycorrhiza.wiki">Mycorrhiza Wiki</a> version:</b> 1.3.0</li>
`)
//line views/stuff.qtpl:225
	if cfg.UseAuth {
//line views/stuff.qtpl:225
		qw422016.N().S(`			<li><b>User count:</b> `)
//line views/stuff.qtpl:226
		qw422016.N().DUL(user.Count())
//line views/stuff.qtpl:226
		qw422016.N().S(`</li>
			<li><b>Home page:</b> <a href="/">`)
//line views/stuff.qtpl:227
		qw422016.E().S(cfg.HomeHypha)
//line views/stuff.qtpl:227
		qw422016.N().S(`</a></li>
			<li><b>Administrators:</b>`)
//line views/stuff.qtpl:228
		for i, username := range user.ListUsersWithGroup("admin") {
//line views/stuff.qtpl:229
			if i > 0 {
//line views/stuff.qtpl:229
				qw422016.N().S(`<span aria-hidden="true">, </span>
`)
//line views/stuff.qtpl:230
			}
//line views/stuff.qtpl:230
			qw422016.N().S(`				<a href="/hypha/`)
//line views/stuff.qtpl:231
			qw422016.E().S(cfg.UserHypha)
//line views/stuff.qtpl:231
			qw422016.N().S(`/`)
//line views/stuff.qtpl:231
			qw422016.E().S(username)
//line views/stuff.qtpl:231
			qw422016.N().S(`">`)
//line views/stuff.qtpl:231
			qw422016.E().S(username)
//line views/stuff.qtpl:231
			qw422016.N().S(`</a>`)
//line views/stuff.qtpl:231
		}
//line views/stuff.qtpl:231
		qw422016.N().S(`</li>
`)
//line views/stuff.qtpl:232
	} else {
//line views/stuff.qtpl:232
		qw422016.N().S(`			<li>This wiki does not use authorization</li>
`)
//line views/stuff.qtpl:234
	}
//line views/stuff.qtpl:234
	qw422016.N().S(`		</ul>
		<p>See <a href="/list">/list</a> for information about hyphae on this wiki.</p>
	</section>
</main>
</div>
`)
//line views/stuff.qtpl:240
}

//line views/stuff.qtpl:240
func WriteAboutHTML(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:240
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:240
	StreamAboutHTML(qw422016)
//line views/stuff.qtpl:240
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:240
}

//line views/stuff.qtpl:240
func AboutHTML() string {
//line views/stuff.qtpl:240
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:240
	WriteAboutHTML(qb422016)
//line views/stuff.qtpl:240
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:240
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:240
	return qs422016
//line views/stuff.qtpl:240
}

//line views/stuff.qtpl:242
func StreamCommonScripts(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:242
	qw422016.N().S(`
`)
//line views/stuff.qtpl:243
	for _, scriptPath := range cfg.CommonScripts {
//line views/stuff.qtpl:243
		qw422016.N().S(`
<script src="`)
//line views/stuff.qtpl:244
		qw422016.E().S(scriptPath)
//line views/stuff.qtpl:244
		qw422016.N().S(`"></script>
`)
//line views/stuff.qtpl:245
	}
//line views/stuff.qtpl:245
	qw422016.N().S(`
`)
//line views/stuff.qtpl:246
}

//line views/stuff.qtpl:246
func WriteCommonScripts(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:246
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:246
	StreamCommonScripts(qw422016)
//line views/stuff.qtpl:246
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:246
}

//line views/stuff.qtpl:246
func CommonScripts() string {
//line views/stuff.qtpl:246
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:246
	WriteCommonScripts(qb422016)
//line views/stuff.qtpl:246
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:246
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:246
	return qs422016
//line views/stuff.qtpl:246
}
