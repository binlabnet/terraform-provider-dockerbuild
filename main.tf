provider "dockerbuild" {
    // TODO - git dir here?
}

data "dockerbuild_git_tree" "sample_service" {
    git_root = path.module
    source_dir = "sample_service"
}