//go:generate go run ../../generate/listpages/main.go -ListOps=DescribeFleets,DescribeImageBuilders,DescribeStacks -Export
//go:generate go run ../../generate/tags/main.go -ListTags -ServiceTagsMap -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package appstream
