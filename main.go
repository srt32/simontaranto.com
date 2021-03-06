package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultRedirectPath = "http://blog.simontaranto.com"
)

func main() {
	http.HandleFunc("/", RedirectHandler)

	port := os.Getenv("port")
	fmt.Printf("Running on port: %v", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Problem in ListenAndServe", err)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	redirectedURL := defaultRedirectPath

	if path, ok := urlMap()[r.URL.Path]; ok {
		redirectedURL += path
	}

	w.Header().Set("Referer", "simontaranto.com")
	w.Header().Set("Referrer", "simontaranto.com")

	http.Redirect(w, r, redirectedURL, 301)
}

func urlMap() map[string]string {
	return map[string]string{
		"/2015/01/11/datetime-parsing-with-go.html":                                    "/post/2015-01-11-datetime-parsing-with-go.html",
		"/2014/11/23/speed-up-javascript-capybara-specs-by-blacklisting-urls.html":     "/post/2014-11-23-speed-up-javascript-capybara-specs-by-blacklisting-urls.html/",
		"/2014/09/16/postgres-window-functions.html":                                   "/post/2014-09-16-postgres-window-functions.html/",
		"/2014/09/06/an-introduction-to-go-tools-and-clean-code.html":                  "/post/2014-09-06-an-introduction-to-go-tools-and-clean-code.html/",
		"/2014/09/06/writing-functional-sql-in-postresql-don-t-do-it.html":             "/post/2014-09-06-writing-functional-sql-in-postresql-don-t-do-it.html/",
		"/2014/08/11/why-postgres-won-t-always-use-an-index.html":                      "/post/2014-08-11-why-postgres-won-t-always-use-an-index.html/",
		"/2014/06/02/workflows-for-writing-migrations-with-rollbacks-in-mind.html":     "/post/2014-06-02-workflows-for-writing-migrations-with-rollbacks-in-mind.html/",
		"/2014/04/16/refactoring-ruby-iteration-patterns-to-the-database.html":         "/post/2014-04-16-refactoring-ruby-iteration-patterns-to-the-database.html/",
		"/2014/03/25/stuck-get-into-the-pairing-mentality.html":                        "/post/2014-03-25-stuck-get-into-the-pairing-mentality.html/",
		"/2014/02/14/n-101-an-example-query-optimization.html":                         "/post/2014-02-14-n-101-an-example-query-optimization.html/",
		"/2014/02/06/writing-and-deploying-multiple-rails-apps-with-nginx.html":        "/post/2014-02-06-writing-and-deploying-multiple-rails-apps-with-nginx.html/",
		"/2014/01/23/doing-more-than-deploying-code-in-a-git-post-receive-hook.html":   "/post/2014-01-23-doing-more-than-deploying-code-in-a-git-post-receive-hook.html/",
		"/2014/01/10/how-activerecord-casts-params-before-validation.html":             "/post/2014-01-10-how-activerecord-casts-params-before-validation.html/",
		"/2013/12/31/using-postgresql-s-tsrange-range-type-with-rails.html":            "/post/2013-12-31-using-postgresql-s-tsrange-range-type-with-rails.html/",
		"/2013/12/11/it-all-comes-together-ruby-js-and-functional-programming.html":    "/post/2013-12-11-it-all-comes-together-ruby-js-and-functional-programming.html/",
		"/2013/12/05/rails_12factor-logging-problems.html":                             "/post/2013-12-05-rails_12factor-logging-problems.html/",
		"/2013/11/17/what-i-learned-setting-up-delayed_paperclip-and-sidekiq.html":     "/post/2013-11-17-what-i-learned-setting-up-delayed_paperclip-and-sidekiq.html/",
		"/2013/11/15/a-refactoring-epiphany.html":                                      "/post/2013-11-15-a-refactoring-epiphany.html/",
		"/2013/10/28/99-bottles-of-robots-our-first-code-retreat.html":                 "/post/2013-10-28-99-bottles-of-robots-our-first-code-retreat.html/",
		"/2013/10/21/array-arithmetic-in-ruby.html":                                    "/post/2013-10-21-array-arithmetic-in-ruby.html/",
		"/2013/10/21/run-a-single-minitest-test.html":                                  "/post/2013-10-21-run-a-single-minitest-test.html/",
		"/2013/10/18/cleaning-up-require-statements-with-ruby-s-global-variable.html":  "/post/2013-10-18-cleaning-up-require-statements-with-ruby-s-global-variable.html/",
		"/2013/10/18/testing-sinatra-apps-with-capybara-and-minitest.html":             "/post/2013-10-18-testing-sinatra-apps-with-capybara-and-minitest.html/",
		"/2013/10/17/bundler-basics-a-primer.html":                                     "/post/2013-10-17-bundler-basics-a-primer.html/",
		"/2013/10/16/set-a-default-rake-task.html":                                     "/post/2013-10-16-set-a-default-rake-task.html/",
		"/2013/10/11/focus-week-reflection.html":                                       "/post/2013-10-11-focus-week-reflection.html/",
		"/2013/10/04/the-power-of-hashes-why-thinking-about-your-problem-matters.html": "/post/2013-10-04-the-power-of-hashes-why-thinking-about-your-problem-matters.html/",
		"/2013/10/02/vim-commands-for-dummies.html":                                    "/post/2013-10-02-vim-commands-for-dummies.html/",
		"/2013/09/20/tests-tmux-and-tears.html":                                        "/post/2013-09-20-tests-tmux-and-tears.html/",
		"/2013/09/18/ruby-enumerable-is-amazing.html":                                  "/post/2013-09-18-ruby-enumerable-is-amazing.html/",
		"/2013/09/13/gschool-week-0.html":                                              "/post/2013-09-13-gschool-week-0.html/",
	}
}
