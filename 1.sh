repo_url=$(git remote get-url origin)
repo_name=$(basename -s .git "$repo_url")
echo "$repo_name"
