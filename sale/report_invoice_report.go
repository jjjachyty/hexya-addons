// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package sale

import (
	"github.com/hexya-erp/hexya/hexya/models"
	"github.com/hexya-erp/hexya/pool"
)

func init() {

	pool.AccountInvoiceReport().AddFields(map[string]models.FieldDefinition{
		"Team": models.Many2OneField{String: "Sales Team", RelationModel: pool.CRMTeam()},
	})

	pool.AccountInvoiceReport().Methods().Select().Extend("",
		func(rs pool.AccountInvoiceReportSet) string {
			return rs.Super().Select() + ", sub.team_id as team_id"
		})

	pool.AccountInvoiceReport().Methods().SubSelect().Extend("",
		func(rs pool.AccountInvoiceReportSet) string {
			return rs.Super().SubSelect() + ", ai.team_id as team_id"
		})

	pool.AccountInvoiceReport().Methods().GroupByClause().Extend("",
		func(rs pool.AccountInvoiceReportSet) string {
			return rs.Super().GroupByClause() + ", ai.team_id"
		})

}
