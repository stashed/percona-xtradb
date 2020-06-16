#!/bin/bash
set -eou pipefail

should_cherry_pick() {
    while IFS=$': \t' read -r -u9 marker v; do
        if [ "$marker" = "/cherry-pick" ]; then
            return 0
        fi
    done 9< <(git show -s --format=%b)
    return 1
}

should_cherry_pick || {
    echo "Skipped cherry picking."
    echo "To automatically cherry pick, add /cherry-pick to commit message body."
    exit 0
}

git remote set-url origin https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git

while IFS=/ read -r -u9 repo branch; do
    git checkout $branch
    pr_branch="master-${GITHUB_SHA:0:8}"${branch#"release"}
    git checkout -b $pr_branch
    git cherry-pick --strategy=recursive -X theirs $GITHUB_SHA
    git push -u origin HEAD -f
    hub pull-request \
        --base $branch \
        --message "[cherry-pick] $(git show -s --format=%s)" \
        --message "$(git show -s --format=%b)" || true
    sleep 2
done 9< <(git branch -r | grep release)
