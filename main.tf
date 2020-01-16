resource "demo_repo_content" "main" {
  organization = "Markieta-Inc"
  repo = "repo1"
  branch = "master"
  file_path = "README.md"
  file_content = <<EOF
# Repository Numba 1

Chris, we really need to be more descriptive with these docs...
EOF
}