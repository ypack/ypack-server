package model

const queryGetPackagesByOS = `
SELECT
  ypack.package.name as pkg_name,
  ypack.package.description as pkg_desc,
  ypack.package.website,
  ypack.version.name as version,
  ypack.version.url,
  ypack.version.checksum,
  ypack.version.os,
  ypack.version.arch,
  ypack.alias.name as alias_name,
  ypack.author.name as author_name,
  ypack.author.contact
FROM 
  package, version, alias, author
WHERE 
  ypack.package.id = ypack.version.package_id AND ypack.package.id = ypack.alias.package_id AND ypack.version.os = ?;
`
