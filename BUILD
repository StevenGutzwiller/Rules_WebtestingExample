load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix
# gazelle:resolve go go github.com/bazelbuild/rules_webtesting/go/webtest @io_bazel_rules_webtesting//go/webtest:webtest.go
gazelle(
    name = "gazelle",
    command = "fix",
    extra_args = [
    ],
)

load("@io_bazel_rules_webtesting//web:go.bzl", "go_web_test_suite")

go_web_test_suite(
    name = "jstestdriver",
    srcs = ["jstestdriver.go"],
    browsers = [
        "@io_bazel_rules_webtesting//browsers:chromium-local",
    ],
    local = True,
    deps = [":go_default_library"],
)
