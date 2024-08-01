import React, { useState } from "react";
import { useMutation, gql } from "@apollo/client";

const ADD_TODO = gql`
  mutation AddTodo($title: String!) {
    addTodo(title: $title) {
      id
      title
      completed
    }
  }
`;

function AddTodo() {
  const [title, setTitle] = useState("");
  const [addTodo] = useMutation(ADD_TODO, {
    refetchQueries: ["GET_TODOS"],
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!title.trim()) return;
    addTodo({ variables: { title } });
    setTitle("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Add a new todo"
      />
      <button type="submit">Add Todo</button>
    </form>
  );
}

export default AddTodo;
