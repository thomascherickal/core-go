// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parse does the parsing stage after lexing
package parse

import (
	"fmt"
	"io"

	"github.com/goki/ki/indent"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
	"github.com/goki/ki/walki"
	"goki.dev/pi/v2/lex"
	"goki.dev/pi/v2/syms"
)

// Ast is a node in the abstract syntax tree generated by the parsing step
// the name of the node (from ki.Node) is the type of the element
// (e.g., expr, stmt, etc)
// These nodes are generated by the parse.Rule's by matching tokens
type Ast struct {
	ki.Node

	// region in source lexical tokens corresponding to this Ast node -- Ch = index in lex lines
	TokReg lex.Reg `desc:"region in source lexical tokens corresponding to this Ast node -- Ch = index in lex lines"`

	// region in source file corresponding to this Ast node
	SrcReg lex.Reg `desc:"region in source file corresponding to this Ast node"`

	// source code corresponding to this Ast node
	Src string `desc:"source code corresponding to this Ast node"`

	// stack of symbols created for this node
	Syms syms.SymStack `desc:"stack of symbols created for this node"`
}

var KiT_Ast = kit.Types.AddType(&Ast{}, AstProps)

// ChildAst returns the Child at given index as an Ast.
// Will panic if index is invalid -- use Try if unsure.
func (ast *Ast) ChildAst(idx int) *Ast {
	return ast.Child(idx).(*Ast)
}

// ChildAstTry returns the child at given index as an Ast -- error if not valid
func (ast *Ast) ChildAstTry(idx int) (*Ast, error) {
	asti, err := ast.ChildTry(idx)
	if err != nil {
		return nil, err
	}
	return asti.(*Ast), nil
}

// ParAst returns the Parent as an Ast.
func (ast *Ast) ParAst() *Ast {
	if ast.Par == nil {
		return nil
	}
	pki := ast.Par.This()
	if pki == nil {
		return nil
	}
	return pki.(*Ast)
}

// NextAst returns the next node in the Ast tree, or nil if none
func (ast *Ast) NextAst() *Ast {
	nxti := walki.Next(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*Ast)
}

// NextSiblingAst returns the next sibling node in the Ast tree, or nil if none
func (ast *Ast) NextSiblingAst() *Ast {
	nxti := walki.NextSibling(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*Ast)
}

// PrevAst returns the previous node in the Ast tree, or nil if none
func (ast *Ast) PrevAst() *Ast {
	nxti := walki.Prev(ast)
	if nxti == nil {
		return nil
	}
	return nxti.(*Ast)
}

// SetTokReg sets the token region for this rule to given region
func (ast *Ast) SetTokReg(reg lex.Reg, src *lex.File) {
	ast.TokReg = reg
	ast.SrcReg = src.TokenSrcReg(ast.TokReg)
	ast.Src = src.RegSrc(ast.SrcReg)
}

// SetTokRegEnd updates the ending token region to given position --
// token regions are typically over-extended and get narrowed as tokens actually match
func (ast *Ast) SetTokRegEnd(pos lex.Pos, src *lex.File) {
	ast.TokReg.Ed = pos
	ast.SrcReg = src.TokenSrcReg(ast.TokReg)
	ast.Src = src.RegSrc(ast.SrcReg)
}

// WriteTree writes the AST tree data to the writer -- not attempting to re-render
// source code -- just for debugging etc
func (ast *Ast) WriteTree(out io.Writer, depth int) {
	ind := indent.Tabs(depth)
	fmt.Fprintf(out, "%v%v: %v\n", ind, ast.Nm, ast.Src)
	for _, k := range ast.Kids {
		ai := k.(*Ast)
		ai.WriteTree(out, depth+1)
	}
}

var AstProps = ki.Props{
	"EnumType:Flag": ki.KiT_Flags,
	"StructViewFields": ki.Props{ // hide in view
		"Flag":  `view:"-"`,
		"Props": `view:"-"`,
	},
	// "CallMethods": ki.PropSlice{
	// 	{"SaveAs", ki.Props{
	// 		"Args": ki.PropSlice{
	// 			{"File Name", ki.Props{
	// 				"default-field": "Filename",
	// 			}},
	// 		},
	// 	}},
	// },
}
