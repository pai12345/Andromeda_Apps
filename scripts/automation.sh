#!/bin/bash

function git_diff() {
  git diff origin/main HEAD --name-only
}

git_diff
