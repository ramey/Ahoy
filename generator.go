package main

type Generator interface {
	Generate()
	Children() []Generator
}

type dirGenerator struct {
	dirname  string
	children []Generator
}

func NewDirGeneraor(dirname string) *dirGenerator {
	return &dirGenerator{
		dirname: dirname,
	}
}

func (dg *dirGenerator) AddFile(file *fileGenerator) *dirGenerator {
	dg.children = append(dg.children, file)
	return dg
}

func (dg *dirGenerator) AddDir(dir *dirGenerator) *dirGenerator {
	dg.children = append(dg.children, dir)
	return dg
}

func (dg *dirGenerator) Generate() {

}

func (dg *dirGenerator) Children() []Generator {
	return dg.children
}

type fileGenerator struct {
	filename string
	code     *codeGenerator
}

func NewFileGenerator(filename string) *fileGenerator {
	return &fileGenerator{
		filename: filename,
	}
}

func (fg *fileGenerator) Generate() {

}

func (fg *fileGenerator) Children() []Generator {
	return []Generator{fg.code}
}

func (fg *fileGenerator) AddCode(code *codeGenerator) *fileGenerator {
	fg.code = code
	return fg
}

type codeGenerator struct {
}

func (cg *codeGenerator) Generate() {

}

func (cg *codeGenerator) Children() []Generator {
	return nil
}
