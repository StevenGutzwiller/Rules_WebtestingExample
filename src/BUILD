load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_webtesting//web:go.bzl", "go_web_test_suite")

go_web_test_suite(
    name = "jstestdriver",
    srcs = ["jstestdriver.go"],
    browsers = [
        "@io_bazel_rules_webtesting//browsers:chromium-local",
    ],
    local = True,
    deps = ["@io_bazel_rules_webtesting//go/webtest:go_default_library"],
)

go_library(
    name = "jstestdriverlib",
    srcs = ["jstestdriver.go"],
    importpath = "src",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_tebeka_selenium//:go_default_library",
        "@io_bazel_rules_webtesting//go/webtest:go_default_library",
    ],
)
