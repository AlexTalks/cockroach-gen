// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

import React from "react";
import _ from "lodash";
import { assert } from "chai";
import { mount } from "enzyme";
import * as sinon from "sinon";

import "src/enzymeInit";
import { SortedTable, ColumnDescriptor } from "src/views/shared/components/sortedtable";
import { SortSetting } from "src/views/shared/components/sortabletable";

class TestRow {
  constructor(public name: string, public value: number) { }
}

const columns: ColumnDescriptor<TestRow>[] = [
  {
    title: "first",
    cell: (tr) => tr.name,
    sort: (tr) => tr.name,
  },
  {
    title: "second",
    cell: (tr) => tr.value.toString(),
    sort: (tr) => tr.value,
    rollup: (trs) => _.sumBy(trs, (tr) => tr.value),
  },
];

class TestSortedTable extends SortedTable<TestRow> {}

function makeTable(
  data: TestRow[], sortSetting?: SortSetting, onChangeSortSetting?: (ss: SortSetting) => void,
) {
  return mount(<TestSortedTable data={data}
                                sortSetting={sortSetting}
                                onChangeSortSetting={onChangeSortSetting}
                                columns={columns}/>);
}

describe("<SortedTable>", function() {
  it("renders the expected table structure.", function() {
    const wrapper = makeTable([new TestRow("test", 1)]);
    assert.lengthOf(wrapper.find("table"), 1, "one table");
    assert.lengthOf(wrapper.find("thead").find("tr"), 1, "one header row");
    assert.lengthOf(wrapper.find("tr.sort-table__row--header"), 1, "column header row");
    assert.lengthOf(wrapper.find("tbody"), 1, "tbody element");
  });

  it("correctly uses onChangeSortSetting", function() {
    const spy = sinon.spy();
    const wrapper = makeTable([new TestRow("test", 1)], undefined, spy);
    wrapper.find("th.sort-table__cell--sortable").first().simulate("click");
    assert.isTrue(spy.calledOnce);
    assert.deepEqual(spy.getCall(0).args[0], {
      sortKey: 0,
      ascending: false,
    } as SortSetting);
  });

  it("correctly sorts data based on sortSetting", function() {
    const data = [
      new TestRow("c", 3),
      new TestRow("d", 4),
      new TestRow("a", 1),
      new TestRow("b", 2),
    ];
    let wrapper = makeTable(data, undefined);
    const assertMatches = (expected: TestRow[]) => {
      const rows = wrapper.find("tbody");
      _.each(expected, (rowData, dataIndex) => {
        const row = rows.childAt(dataIndex);
        assert.equal(row.childAt(0).text(), rowData.name, "first columns match");
        assert.equal(row.childAt(1).text(), rowData.value.toString(), "second columns match");
      });
    };
    assertMatches(data);
    wrapper = makeTable(data, {sortKey: 0, ascending: true});
    assertMatches(_.sortBy(data, (r) => r.name));
    wrapper.setProps({uiSortSetting: {sortKey: 1, ascending: true} as SortSetting});
    assertMatches(_.sortBy(data, (r) => r.value));
  });
});
