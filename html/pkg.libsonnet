local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-libs',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'html',
  path: 'html',
  target: 'h',
}, |||
  Jsonnet library for generating HTML.
|||)
