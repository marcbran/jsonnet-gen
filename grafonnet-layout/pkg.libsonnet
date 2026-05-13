local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-libs',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'grafonnet-layout',
  path: 'grafonnet-layout',
  target: 'lt',
}, |||
  Jsonnet library that adds simple layout functions to grafonnet.
|||)
