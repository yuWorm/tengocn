// parser/syntax_extension_test.go - 新增测试文件

package parser_test

import (
	"testing"

	. "github.com/d5/tengo/v2/parser"
)

func TestParseHasValue(t *testing.T) {
	// 测试简单的 has_value 表达式
	expectParse(t, "如果 x 有值 {}", func(p pfn) []Stmt {
		return stmts(
			ifStmt(
				nil,
				hasValueExpr(ident("x", p(1, 4)), p(1, 6)),
				blockStmt(p(1, 16), p(1, 17)),
				nil,
				p(1, 1)))
	})
}

func TestParseIs(t *testing.T) {
	// 测试 is True
	expectParse(t, "if x is True {}", func(p pfn) []Stmt {
		return stmts(
			ifStmt(
				nil,
				isExpr(ident("x", p(1, 4)), ident("True", p(1, 9)), p(1, 6)),
				blockStmt(p(1, 14), p(1, 15)),
				nil,
				p(1, 1)))
	})

	// 测试 is 数值
	expectParse(t, "if y is 0 {}", func(p pfn) []Stmt {
		return stmts(
			ifStmt(
				nil,
				isExpr(ident("y", p(1, 4)), intLit(0, p(1, 9)), p(1, 6)),
				blockStmt(p(1, 11), p(1, 12)),
				nil,
				p(1, 1)))
	})
}

// 辅助构造函数
func hasValueExpr(expr Expr, pos Pos) *HasValueExpr {
	return &HasValueExpr{
		Expr:        expr,
		HasValuePos: pos,
	}
}

func isExpr(lhs, rhs Expr, pos Pos) *IsExpr {
	return &IsExpr{
		LHS:   lhs,
		RHS:   rhs,
		IsPos: pos,
	}
}
