package db

type QueryType map[string]string

type QueryMap map[string]QueryType

var queryMap = QueryMap{
	"sqlite3":  sqliteQueries,
	"postgres": postgresQueries,
}
var sqliteQueries = QueryType{
	"getUserByID":      "SELECT id, name, username FROM users WHERE id=$1",
	"getAllUsers":      "SELECT id, name, username FROM users",
	"createUser":       "INSERT INTO users (name, username) VALUES ($1, $2) RETURNING id",
	"updateUser":       "UPDATE users SET name=$1, username=$2 WHERE id=$3",
	"deleteUser":       "DELETE FROM users WHERE id=$1",
	"getTodoByID":      "SELECT id, title, completed, user_id FROM todos WHERE id=$1",
	"getTodosByUserID": "SELECT id, title, completed, user_id FROM todos WHERE user_id=$1",
	"createTodo":       "INSERT INTO todos (title, completed, user_id) VALUES ($1, $2, $3) RETURNING id",
	"updateTodo":       "UPDATE todos SET title=$1,	 completed=$2, user_id=$3 WHERE id=$4",
	"deleteTodo":       "DELETE FROM todos WHERE id=$1",
}
var postgresQueries = QueryType{
	"getUserByID":      "SELECT id, name, username FROM users WHERE id=$1",
	"getAllUsers":      "SELECT id, name, username FROM users",
	"createUser":       "INSERT INTO users (name, username) VALUES ($1, $2) RETURNING id",
	"updateUser":       "UPDATE users SET name=$1, username=$2 WHERE id=$3",
	"deleteUser":       "DELETE FROM users WHERE id=$1",
	"getTodoByID":      "SELECT id, title, completed, user_id FROM todos WHERE id=$1",
	"getTodosByUserID": "SELECT id, title, completed, user_id FROM todos WHERE user_id=$1",
	"createTodo":       "INSERT INTO todos (title, completed, user_id) VALUES ($1, $2, $3) RETURNING id",
	"updateTodo":       "UPDATE todos SET title=$1,	 completed=$2, user_id=$3 WHERE id=$4",
	"deleteTodo":       "DELETE FROM todos WHERE id=$1",
}

func GetQuery(dbType, queryName string) string {
	if queries, ok := queryMap[dbType]; ok {
		if query, ok := queries[queryName]; ok {
			return query
		}
	}
	return ""
}
