import React from "react";
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import TodoList from "./components/TodoList";
import AddTodo from "./components/AddTodo";

const client = new ApolloClient({
  uri: "http://localhost:8080/graphql",
  cache: new InMemoryCache(),
});

function App() {
  return (
    <ApolloProvider client={client}>
      <div className="App">
        <h1>Todo App</h1>
        <AddTodo />
        <TodoList />
      </div>
    </ApolloProvider>
  );
}

export default App;
