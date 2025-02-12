package validate

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

func TestArrayForEachValidator(t *testing.T) {
	t.Parallel()

	rootPath := path.Root("test")

	type testCase struct {
		val         tftypes.Value
		f           func(context.Context, tftypes.Value) (attr.Value, error)
		validator   tfsdk.AttributeValidator
		expectError bool
	}
	tests := map[string]testCase{
		"not a list": {
			val:         tftypes.NewValue(tftypes.Bool, true),
			f:           types.BoolType.ValueFromTerraform,
			validator:   StringInSlice([]string{"alpha", "beta", "gamma"}),
			expectError: true,
		},

		"unknown list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"null list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, nil),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"empty list": {
			val:       tftypes.NewValue(tftypes.List{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:         types.ListType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"valid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:         types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"invalid list of string": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "delta"),
			}),
			f:           types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			validator:   StringInSlice([]string{"alpha", "beta", "gamma"}),
			expectError: true,
		},
		"list of string with unknown element": {
			val: tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:         types.ListType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},

		"unknown set": {
			val:       tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, tftypes.UnknownValue),
			f:         types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"null set": {
			val:       tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, nil),
			f:         types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"empty set": {
			val:       tftypes.NewValue(tftypes.Set{ElementType: tftypes.Number}, []tftypes.Value{}),
			f:         types.SetType{ElemType: types.NumberType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"valid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:         types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
		"invalid set of string": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, "beta"),
				tftypes.NewValue(tftypes.String, "gamma"),
				tftypes.NewValue(tftypes.String, "delta"),
			}),
			f:           types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			validator:   StringInSlice([]string{"alpha", "beta", "gamma"}),
			expectError: true,
		},
		"set of string with unknown element": {
			val: tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "alpha"),
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				tftypes.NewValue(tftypes.String, "gamma"),
			}),
			f:         types.SetType{ElemType: types.StringType}.ValueFromTerraform,
			validator: StringInSlice([]string{"alpha", "beta", "gamma"}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			val, err := test.f(ctx, test.val)

			if err != nil {
				t.Fatalf("got unexpected error: %s", err)
			}

			request := tfsdk.ValidateAttributeRequest{
				AttributePath:   rootPath,
				AttributeConfig: val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			ArrayForEach(test.validator).Validate(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", tfresource.DiagsError(response.Diagnostics))
			}
		})
	}
}

func printDiagnostics(ds diag.Diagnostics) string {
	s := make([]string, len(ds))
	for i, d := range ds {
		s[i] = printDiagnostic(d)
	}

	return "[" + strings.Join(s, ", ") + "]"
}

func printDiagnostic(d diag.Diagnostic) string {
	b := &strings.Builder{}
	fmt.Fprintf(b, `%s "%s": "%s"`, d.Severity().String(), d.Summary(), d.Detail())

	if v, ok := d.(diag.DiagnosticWithPath); ok {
		fmt.Fprintf(b, " at %s", v.Path().String())
	}

	return b.String()
}
