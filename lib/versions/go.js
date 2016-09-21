const findVersions = require('find-versions');
const shell = require('shelljs');

const output = shell.exec('go version', { silent: true }).stdout;
const version = findVersions(output, {loose: true})[0];

module.exports = version;