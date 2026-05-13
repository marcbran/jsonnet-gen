local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-libs',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'grafonnet-query',
  path: 'grafonnet-query',
  target: 'q',
}, |||
  Jsonnet library that adds shortcuts to grafonnet for working with Jsonnet-based query languages.
|||)
