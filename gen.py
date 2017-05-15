import os

data = open("public_suffix_list.txt").readlines()
data = map(lambda s:s.strip(), data)

def lineFilter(s):
    if not s:
        return False
    if s.startswith("//") or s.startswith("!"):
        return False
    return True
data = filter(lineFilter, data)
data = map(lambda s: s if not s.startswith("*.") else s[2:], data)

root = dict()
for item in data:
    tokens = item.split(".")
    tokens.reverse()
    p = root
    for t in tokens:
        if t not in p:
            p[t] = dict()
        p = p[t]

def genTree(name, ch, dep):
    if not ch:
        return "    " * dep + 'node{name:"%s"}' % name
    children = list()
    for n, c in sorted(ch.items()):
        children.append(genTree(n, c, dep+2) + ",")
    ident = " " * 4 * (dep + 1)
    fmt = " " * 4 * dep + """node{
%(ident)sname:"%(name)s",
%(ident)sch: []node{
%(ch)s
%(ident)s}}"""
    return fmt % dict(name=name, ident=ident, ch="\n".join(children))

tmpl = """package maindomain
var PublicSuffixTree = %s 
""" 
code = open("public_suffix_tree.go", "w")
code.write(tmpl % genTree("ROOT", root, 0))
code.close()

os.system("go fmt public_suffix_tree.go")
