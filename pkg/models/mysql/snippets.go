package mysql
import (
	"database/sql"
	"main/pkg/models"
)
type SnippetModel struct{
  DB *sql.DB
}
func(m *SnippetModel)Insert(Title , Content string) (int, error){
  stmt := `INSERT INTO tarefas (title, content) VALUES(?,?)`

  result, err := m.DB.Exec(stmt, Title, Content)
  if err != nil{
    return 0, err
  }
  
  id, err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  return int(id),nil
}

func(m *SnippetModel) Get(id int)(*models.Tarefas, error){
  stmt := `SELECT id, title, content FROM tarefas WHERE id = ?`
  row := m.DB.QueryRow(stmt, id)

  s := &models.Tarefas{}

  err := row.Scan(&s.ID, &s.Title, &s.Content)
  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  }else if err != nil{
    return nil, err
  }
  
  return s, nil
}

func(m * SnippetModel) Latest()([]*models.Tarefas, error){
  stmt := `SELECT id, title, content FROM tarefas ORDER BY id DESC`

  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  snippets := []*models.Tarefas{}
  for rows.Next(){
    s := &models.Tarefas{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content)
    if err != nil{
      return nil, err
    }
    snippets = append(snippets, s)
  }
  err = rows.Err()
  if err != nil {
    return  nil, err
  }
  return snippets, nil
}

func(m *SnippetModel) Delete(ID int) (int, error){
  stmt := `DELETE FROM tarefas WHERE id = ?`

  _, err := m.DB.Exec(stmt, ID)
  if err != nil{
    return 0, err
  }
  
  return 0, nil
}

func(m *SnippetModel) Edit(ID int, Title, Content string) (int, error){
  stmt := `UPDATE tarefas SET title = ?, content = ? WHERE id = ?`

  _, err := m.DB.Exec(stmt, Title, Content, ID)
  if err != nil{
    return 0, err
  }
  
  return 0, nil
}