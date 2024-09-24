package engine

// for _, name := range msg.data.Columns {
// 	columns = append(columns, table.Column{Title: name, Width: 10})
// }
//
// rows := []table.Row{}
// for rowIdx := 0; rowIdx < len(msg.data.Data[columns[0].Title]); rowIdx++ {
// 	row := make(table.Row, 0, len(msg.data.Data[columns[0].Title]))
// 	for _, col := range msg.data.Data {
// 		row = append(row, fmt.Sprintf("%v", col[rowIdx]))
// 	}
//
// 	rows = append(rows, row)
// }
//
// t := table.New(
// 	table.WithColumns(columns),
// 	table.WithRows(rows),
// 	table.WithHeight(20),
// )
// s := table.DefaultStyles()
//
// s.Header = s.Header.
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderForeground(lipgloss.Color("240")).
// 	BorderBottom(true).
// 	Bold(false)
// s.Selected = s.Selected.
// 	Foreground(lipgloss.Color("240")).
// 	Background(lipgloss.Color("0")).
// 	Bold(false)
// t.SetStyles(s)
