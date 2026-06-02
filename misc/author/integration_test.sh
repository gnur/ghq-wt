#!/bin/sh
set -e

tmpdir=$(mktemp -d)

cleanup() {
    code=$?
    rm -rf $tmpdir
    exit $code
}
trap cleanup EXIT

set -x

export GHQ_ROOT=$tmpdir

: testing 'ghq get'
    ghq get gnur/ghq-wt
    ghq get --bare x-motemen/gore
    ghq get --partial blobless x-motemen/blogsync
    ghq get --partial treeless x-motemen/gobump

    test -d $tmpdir/github.com/gnur/ghq-wt/.bare
    test -d $tmpdir/github.com/x-motemen/gore.git/refs
    grep --quiet "partialclonefilter = blob:none" $tmpdir/github.com/x-motemen/blogsync/.git/config
    grep --quiet "partialclonefilter = tree:0" $tmpdir/github.com/x-motemen/gobump/.git/config

: testing 'ghq list'
    cat <<EOF | sort > $tmpdir/expect
github.com/x-motemen/blogsync
github.com/gnur/ghq-wt/master
github.com/x-motemen/gobump
github.com/x-motemen/gore.git
EOF
    ghq list | sort > $tmpdir/got
    diff -u $tmpdir/expect $tmpdir/got

: testing 'input | ghq get -u'
    ghq list | ghq get -u

: testing 'ghq create'
    test "$(ghq create Songmu/hoge)" = "$tmpdir/github.com/Songmu/hoge"
    test -d $tmpdir/github.com/Songmu/hoge/.git
