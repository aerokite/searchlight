#!/usr/bin/env python


# http://stackoverflow.com/a/14050282
def check_antipackage():
    from sys import version_info
    sys_version = version_info[:2]
    found = True
    if sys_version < (3, 0):
        # 'python 2'
        from pkgutil import find_loader
        found = find_loader('antipackage') is not None
    elif sys_version <= (3, 3):
        # 'python <= 3.3'
        from importlib import find_loader
        found = find_loader('antipackage') is not None
    else:
        # 'python >= 3.4'
        from importlib import util
        found = util.find_spec('antipackage') is not None
    if not found:
        print('Install missing package "antipackage"')
        print('Example: pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage')
        from sys import exit
        exit(1)
check_antipackage()

# ref: https://github.com/ellisonbg/antipackage
import antipackage
from github.appscode.libbuild import libbuild

import os
import os.path
import subprocess
import sys
from os.path import expandvars

libbuild.REPO_ROOT = expandvars('$GOPATH') + '/src/github.com/appscode/searchlight'
BUILD_METADATA = libbuild.metadata(libbuild.REPO_ROOT)
libbuild.BIN_MATRIX = {
    'searchlight': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'hostfacts': {
        'type': 'go',
        'go_version': False,
        'release': True,
        'use_cgo': False,
        'distro': {
            'linux': ['amd64'],
        }
    },
    'hyperalert': {
        'type': 'go',
        'go_version': True,
        'release': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_component_status': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_influx_query': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_json_path': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_node_count': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_node_status': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_pod_exists': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_pod_status': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_prometheus_metric': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_volume': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_kube_event': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'check_kube_exec': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'notifier': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'linux': ['amd64']
        }
    },
    'hello_icinga': {
        'type': 'go',
        'go_version': True,
        'release': True,
        'distro': {
            'linux': ['amd64']
        }
    },
}
libbuild.BUCKET_MATRIX = {
    'prod': 'gs://appscode-cdn',
    'dev': 'gs://appscode-dev'
}


def call(cmd, stdin=None, cwd=libbuild.REPO_ROOT):
    print(cmd)
    return subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def die(status):
    if status:
        sys.exit(status)


def check_output(cmd, stdin=None, cwd=libbuild.REPO_ROOT):
    print(cmd)
    return subprocess.check_output([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def version():
    # json.dump(BUILD_METADATA, sys.stdout, sort_keys=True, indent=2)
    for k in sorted(BUILD_METADATA):
        print(k + '=' + BUILD_METADATA[k])


def fmt():
    libbuild.ungroup_go_imports('cmd', 'data', 'pkg', 'plugins', 'util')
    die(call('goimports -w cmd data pkg plugins util'))
    call('gofmt -s -w cmd data pkg plugins util')


def vet():
    call('go vet ./cmd/... ./data/... ./pkg/... ./plugins/... ./util/...')


def lint():
    call('golint ./cmd/...')
    call('golint ./data/...')
    call('golint ./pkg/...')
    call('golint ./plugins/...')
    call('golint ./util/...')


def gen_assets():
    die(call('go-bindata -ignore=\\.go -ignore=\\.DS_Store -mode=0644 -modtime=1453795200 -o bindata.go -pkg files ./...', cwd=libbuild.REPO_ROOT + '/data/files'))


def gen_extpoints():
    die(call('go generate cmd/notifier/main.go'))


def gen():
    gen_assets()
    gen_extpoints()
    fmt()


def build_cmd(name):
    cfg = libbuild.BIN_MATRIX[name]
    if cfg['type'] == 'go':
        if 'distro' in cfg:
            for goos, archs in cfg['distro'].items():
                for goarch in archs:
                    libbuild.go_build(name, goos, goarch, main='cmd/{}/*.go'.format(name))
        else:
            libbuild.go_build(name, libbuild.GOHOSTOS, libbuild.GOHOSTARCH, main='cmd/{}/*.go'.format(name))


def build_cmds():
    gen()
    for name in libbuild.BIN_MATRIX:
        build_cmd(name)


def build(name=None):
    if name:
        cfg = libbuild.BIN_MATRIX[name]
        if cfg['type'] == 'go':
            gen()
            build_cmd(name)
    else:
        build_cmds()


def push(name=None):
    if name:
        bindir = libbuild.REPO_ROOT + '/dist/' + name
        push_bin(bindir)
    else:
        dist = libbuild.REPO_ROOT + '/dist'
        for name in os.listdir(dist):
            d = dist + '/' + name
            if os.path.isdir(d):
                push_bin(d)


def push_bin(bindir):
    call('rm -f *.md5', cwd=bindir)
    call('rm -f *.sha1', cwd=bindir)
    for f in os.listdir(bindir):
        if os.path.isfile(bindir + '/' + f):
            libbuild.upload_to_cloud(bindir, f, BUILD_METADATA['version'])


def update_registry():
    libbuild.update_registry(BUILD_METADATA['version'])


def install():
    die(call('GO15VENDOREXPERIMENT=1 ' + libbuild.GOC + ' install ./cmd/...'))


def default():
    gen()
    die(call('GO15VENDOREXPERIMENT=1 ' + libbuild.GOC + ' install ./cmd/...'))


if __name__ == "__main__":
    if len(sys.argv) > 1:
        # http://stackoverflow.com/a/834451
        # http://stackoverflow.com/a/817296
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        default()
