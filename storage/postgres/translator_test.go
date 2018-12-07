/*
 * Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package postgres

import (
	"fmt"

	"github.com/Peripli/service-manager/pkg/selection"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Postgres Translator", func() {

	Describe("translate list", func() {

		baseQuery, baseTableName, labelsTableName := "", "testTable", "testLabelTable"
		var criteria []selection.Criterion

		Context("No query", func() {
			It("Should return base query", func() {
				actualQuery, actualQueryParams, err := buildListQueryWithParams(baseQuery, baseTableName, labelsTableName, criteria)
				Expect(err).ToNot(HaveOccurred())
				Expect(actualQuery).To(Equal(baseQuery))
				Expect(actualQueryParams).To(BeEmpty())
			})
		})

		Context("Label query", func() {
			Context("Called with valid input", func() {
				It("Should return proper result", func() {
					criteria = []selection.Criterion{
						{
							LeftOp:   "orgId",
							Operator: selection.InOperator,
							RightOp:  []string{"o1", "o2", "o3"},
							Type:     selection.LabelQuery,
						},
						{
							LeftOp:   "clusterId",
							Operator: selection.InOperator,
							RightOp:  []string{"c1", "c2"},
							Type:     selection.LabelQuery,
						},
					}
					actualQuery, actualQueryParams, err := buildListQueryWithParams(baseQuery, baseTableName, labelsTableName, criteria)
					Expect(err).ToNot(HaveOccurred())
					Expect(actualQuery).To(Equal(fmt.Sprintf(" WHERE %[1]s.key = ? AND %[1]s.val IN (?, ?, ?) AND %[1]s.key = ? AND %[1]s.val IN (?, ?);", labelsTableName)))

					expectedQueryParams := buildExpectedQueryParams(criteria)
					Expect(actualQueryParams).To(Equal(expectedQueryParams))
				})
			})

			Context("Called with multivalue operator and single value", func() {
				It("Should return proper result surrounded in brackets", func() {
					criteria = []selection.Criterion{
						{
							LeftOp:   "orgId",
							Operator: selection.InOperator,
							RightOp:  []string{"o1"},
							Type:     selection.LabelQuery,
						},
					}
					actualQuery, actualQueryParams, err := buildListQueryWithParams(baseQuery, baseTableName, labelsTableName, criteria)
					Expect(err).ToNot(HaveOccurred())
					Expect(actualQuery).To(Equal(fmt.Sprintf(" WHERE %[1]s.key = ? AND %[1]s.val IN (?);", labelsTableName)))

					expectedQueryParams := buildExpectedQueryParams(criteria)
					Expect(actualQueryParams).To(Equal(expectedQueryParams))
				})
			})
		})
		Context("Field query", func() {
			Context("Called with valid input", func() {
				It("Should return proper result", func() {
					criteria = []selection.Criterion{
						{
							LeftOp:   "platformId",
							Operator: selection.EqualsOperator,
							RightOp:  []string{"5"},
							Type:     selection.FieldQuery,
						},
					}
					actualQuery, actualQueryParams, err := buildListQueryWithParams(baseQuery, baseTableName, labelsTableName, criteria)
					Expect(err).ToNot(HaveOccurred())
					Expect(actualQuery).To(Equal(fmt.Sprintf(" WHERE %s.%s %s ?;", baseTableName, criteria[0].LeftOp, criteria[0].Operator)))

					expectedQueryParams := buildExpectedQueryParams(criteria)
					Expect(actualQueryParams).To(Equal(expectedQueryParams))
				})
			})

			Context("Called with multivalue operator and single value", func() {
				It("Should return proper result surrounded in brackets", func() {
					criteria = []selection.Criterion{
						{
							LeftOp:   "platformId",
							Operator: selection.InOperator,
							RightOp:  []string{"1"},
							Type:     selection.FieldQuery,
						},
					}
					actualQuery, actualQueryParams, err := buildListQueryWithParams(baseQuery, baseTableName, labelsTableName, criteria)
					Expect(err).ToNot(HaveOccurred())
					Expect(actualQuery).To(Equal(fmt.Sprintf(" WHERE %s.%s %s (?);", baseTableName, criteria[0].LeftOp, criteria[0].Operator)))

					expectedQueryParams := buildExpectedQueryParams(criteria)
					Expect(actualQueryParams).To(Equal(expectedQueryParams))
				})
			})

		})
	})
})

func buildExpectedQueryParams(criteria []selection.Criterion) interface{} {
	var expectedQueryParams []interface{}
	for _, criterion := range criteria {
		if criterion.Type == selection.LabelQuery {
			expectedQueryParams = append(expectedQueryParams, criterion.LeftOp)
		}
		for _, param := range criterion.RightOp {
			expectedQueryParams = append(expectedQueryParams, param)
		}
	}
	return expectedQueryParams
}