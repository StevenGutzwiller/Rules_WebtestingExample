load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    command = "fix",
    extra_args = [
        "src/",
    ],
)

load("@io_bazel_rules_webtesting//web:py.bzl", "py_web_test_suite")

py_web_test_suite(
    name = "jstestdriver",
    srcs = ["jstestdriver.go"],
    browsers = [
        "@io_bazel_rules_webtesting//browsers:chromium-local",
    ],
    local = True,
    deps = ["@io_bazel_rules_webtesting//testing/web"],
)