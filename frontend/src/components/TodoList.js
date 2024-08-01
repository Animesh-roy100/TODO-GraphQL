import React from "react";
import { useQuery, gql } from "@apollo/client";

const GET_TODOS = gql`
  query {
    todos {
      id
      title
      completed
    }
  }
`;

function TodoList() {
  const { loading, error, data } = useQuery(GET_TODOS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  return (
    <ul>
      {data.todos.map((todo) => (
        <li key={todo.id}>
          {todo.title} - {todo.completed ? "Completed" : "Not completed"}
        </li>
      ))}
    </ul>
  );
}

export default TodoList;
