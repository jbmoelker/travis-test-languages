#!/usr/bin/env python

from jinja2 import Template
template = Template('Hello from a {{ lang }} script!')
output = template.render(lang='Python')

print(output)