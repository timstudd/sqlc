// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/timstudd/sqlc/internal/sql/ast"
	"github.com/timstudd/sqlc/internal/sql/catalog"
)

func Seg() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = []*catalog.Function{
		{
			Name: "seg_center",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "seg_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "seg_contained",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_contains",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_different",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_ge",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_gt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_in",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "cstring"},
				},
			},
			ReturnType: &ast.TypeName{Name: "seg"},
		},
		{
			Name: "seg_inter",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "seg"},
		},
		{
			Name: "seg_le",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_lower",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "seg_lt",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_out",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "cstring"},
		},
		{
			Name: "seg_over_left",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_over_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_overlap",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_right",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_same",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "boolean"},
		},
		{
			Name: "seg_size",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
		{
			Name: "seg_union",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "seg"},
		},
		{
			Name: "seg_upper",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "seg"},
				},
			},
			ReturnType: &ast.TypeName{Name: "real"},
		},
	}
	return s
}
