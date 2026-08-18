import {
  useEffect,
  useState,
} from "react";

import Navbar from "../components/Navbar";
import TodoForm from "../components/TodoForm";
import TodoItem from "../components/TodoItem";

import {
  getTodos,
  createTodo,
  updateTodo,
  deleteTodo,
} from "../api/todoApi";

function Dashboard() {

  const [todos, setTodos] =
    useState([]);

  const [loading, setLoading] =
    useState(true);

  const [error, setError] =
    useState("");

  /*
   * Load todos
   */

  useEffect(() => {
    loadTodos();
  }, []);

  const loadTodos = async () => {

    try {

      setLoading(true);

      setError("");

      const data =
        await getTodos();

      setTodos(data);

    } catch (error) {

      console.error(
        "Load todos error:",
        error
      );

      setError(
        error.response?.data?.error ||
        "Failed to load todos"
      );

    } finally {

      setLoading(false);

    }
  };

  /*
   * CREATE
   */

  const handleCreateTodo =
    async ({
      title,
      description,
    }) => {

      const newTodo =
        await createTodo(
          title,
          description
        );

      setTodos(
        (currentTodos) => [
          newTodo,
          ...currentTodos,
        ]
      );
    };

  /*
   * TOGGLE COMPLETED
   */

  const handleToggle =
    async (todo) => {

      try {

        setError("");

        const updatedTodo =
          await updateTodo(
            todo.id,
            {
              completed:
                !todo.completed,
            }
          );

        setTodos(
          (currentTodos) =>
            currentTodos.map(
              (item) =>
                item.id ===
                updatedTodo.id
                  ? updatedTodo
                  : item
            )
        );

      } catch (error) {

        console.error(
          "Toggle todo error:",
          error
        );

        setError(
          error.response?.data?.error ||
          "Failed to update todo"
        );
      }
    };

  /*
   * EDIT
   */

  const handleEdit =
    async (todo) => {

      const newTitle =
        window.prompt(
          "Enter new title:",
          todo.title
        );

      if (
        newTitle === null
      ) {
        return;
      }

      if (
        !newTitle.trim()
      ) {
        return;
      }

      const newDescription =
        window.prompt(
          "Enter new description:",
          todo.description ||
            ""
        );

      if (
        newDescription === null
      ) {
        return;
      }

      try {

        setError("");

        const updatedTodo =
          await updateTodo(
            todo.id,
            {
              title:
                newTitle.trim(),

              description:
                newDescription.trim(),
            }
          );

        setTodos(
          (currentTodos) =>
            currentTodos.map(
              (item) =>
                item.id ===
                updatedTodo.id
                  ? updatedTodo
                  : item
            )
        );

      } catch (error) {

        console.error(
          "Edit todo error:",
          error
        );

        setError(
          error.response?.data?.error ||
          "Failed to edit todo"
        );
      }
    };

  /*
   * DELETE
   */

  const handleDelete =
    async (id) => {

      const confirmed =
        window.confirm(
          "Are you sure you want to delete this todo?"
        );

      if (!confirmed) {
        return;
      }

      try {

        setError("");

        await deleteTodo(id);

        setTodos(
          (currentTodos) =>
            currentTodos.filter(
              (todo) =>
                todo.id !== id
            )
        );

      } catch (error) {

        console.error(
          "Delete todo error:",
          error
        );

        setError(
          error.response?.data?.error ||
          "Failed to delete todo"
        );
      }
    };

  return (
    <div className="dashboard">

      <Navbar />

      <main className="dashboard-content">

        <TodoForm
          onTodoCreated={
            handleCreateTodo
          }
        />

        <section className="todos-section">

          <div className="section-header">

            <h2>
              My Tasks
            </h2>

            <span>
              {todos.length}{" "}
              {todos.length === 1
                ? "task"
                : "tasks"}
            </span>

          </div>

          {error && (
            <div className="error">
              {error}
            </div>
          )}

          {loading ? (

            <div className="loading">
              Loading todos...
            </div>

          ) : todos.length === 0 ? (

            <div className="empty-state">

              <h3>
                No todos yet
              </h3>

              <p>
                Create your first
                task above.
              </p>

            </div>

          ) : (

            <div className="todo-list">

              {todos.map(
                (todo) => (

                  <TodoItem
                    key={todo.id}
                    todo={todo}
                    onToggle={
                      handleToggle
                    }
                    onEdit={
                      handleEdit
                    }
                    onDelete={
                      handleDelete
                    }
                  />

                )
              )}

            </div>

          )}

        </section>

      </main>

    </div>
  );
}

export default Dashboard;