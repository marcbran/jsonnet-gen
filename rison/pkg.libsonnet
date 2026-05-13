local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-libs',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'rison',
  path: 'rison',
  target: 'rison',
}, |||
  Jsonnet library that implements conversion to rison.
|||)
