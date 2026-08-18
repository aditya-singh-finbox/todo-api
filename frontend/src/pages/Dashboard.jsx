import {
  useEffect,
  useMemo,
  useState,
} from "react";

import Navbar from "../components/Navbar";
import TodoForm from "../components/TodoForm";
import TodoItem from "../components/TodoItem";
import EditTodoModal from "../components/EditTodoModal";
import DeleteTodoModal from "../components/DeleteTodoModal";

import {
  getTodos,
  createTodo,
  updateTodo,
  deleteTodo,
} from "../api/todoApi";

function Dashboard() {
  // =========================
  // Todo State
  // =========================

  const [todos, setTodos] = useState([]);

  const [loading, setLoading] =
    useState(true);

  const [error, setError] =
    useState("");

  // =========================
  // Search & Filter State
  // =========================

  const [search, setSearch] =
    useState("");

  const [filter, setFilter] =
    useState("all");

  // =========================
  // Modal State
  // =========================

  const [editingTodo, setEditingTodo] =
    useState(null);

  const [deletingTodo, setDeletingTodo] =
    useState(null);

  const [deleteLoading, setDeleteLoading] =
    useState(false);

  // =========================
  // Load Todos
  // =========================

  useEffect(() => {
    loadTodos();
  }, []);

  const loadTodos = async () => {
    try {
      setLoading(true);
      setError("");

      const data = await getTodos();

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

  // =========================
  // Create Todo
  // =========================

  const handleCreateTodo = async ({
    title,
    description,
  }) => {
    try {
      setError("");

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
    } catch (error) {
      console.error(
        "Create todo error:",
        error
      );

      setError(
        error.response?.data?.error ||
          "Failed to create todo"
      );

      throw error;
    }
  };

  // =========================
  // Toggle Todo
  // =========================

  const handleToggle = async (
    todo
  ) => {
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

  // =========================
  // Open Edit Modal
  // =========================

  const handleEdit = (todo) => {
    setEditingTodo(todo);
  };

  // =========================
  // Save Edited Todo
  // =========================

  const handleSaveEdit = async ({
    title,
    description,
  }) => {
    if (!editingTodo) {
      return;
    }

    try {
      setError("");

      const updatedTodo =
        await updateTodo(
          editingTodo.id,
          {
            title,
            description,
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

      setEditingTodo(null);
    } catch (error) {
      console.error(
        "Edit todo error:",
        error
      );

      setError(
        error.response?.data?.error ||
          "Failed to update todo"
      );

      throw error;
    }
  };

  // =========================
  // Open Delete Modal
  // =========================

  const handleDelete = (todo) => {
    setDeletingTodo(todo);
  };

  // =========================
  // Confirm Delete
  // =========================

  const handleConfirmDelete =
    async () => {
      if (!deletingTodo) {
        return;
      }

      try {
        setDeleteLoading(true);
        setError("");

        await deleteTodo(
          deletingTodo.id
        );

        setTodos(
          (currentTodos) =>
            currentTodos.filter(
              (todo) =>
                todo.id !==
                deletingTodo.id
            )
        );

        setDeletingTodo(null);
      } catch (error) {
        console.error(
          "Delete todo error:",
          error
        );

        setError(
          error.response?.data?.error ||
            "Failed to delete todo"
        );
      } finally {
        setDeleteLoading(false);
      }
    };

  // =========================
  // Statistics
  // =========================

  const totalTodos =
    todos.length;

  const completedTodos =
    todos.filter(
      (todo) => todo.completed
    ).length;

  const activeTodos =
    totalTodos -
    completedTodos;

  // =========================
  // Search + Filter
  // =========================

  const filteredTodos =
    useMemo(() => {
      return todos.filter(
        (todo) => {
          // Filter by status

          if (
            filter === "active" &&
            todo.completed
          ) {
            return false;
          }

          if (
            filter === "completed" &&
            !todo.completed
          ) {
            return false;
          }

          // Search

          const searchText =
            search
              .toLowerCase()
              .trim();

          if (!searchText) {
            return true;
          }

          const title =
            todo.title
              ?.toLowerCase() || "";

          const description =
            todo.description
              ?.toLowerCase() || "";

          return (
            title.includes(
              searchText
            ) ||
            description.includes(
              searchText
            )
          );
        }
      );
    }, [todos, filter, search]);

  // =========================
  // Render
  // =========================

  return (
    <div className="dashboard">

      {/* =========================
          Navbar
          ========================= */}

      <Navbar />

      <main className="dashboard-content">

        {/* =========================
            Create Todo
            ========================= */}

        <TodoForm
          onTodoCreated={
            handleCreateTodo
          }
        />

        {/* =========================
            Statistics
            ========================= */}

        <div className="stats-grid">

          <div className="stat-card">

            <span>
              Total
            </span>

            <strong>
              {totalTodos}
            </strong>

          </div>

          <div className="stat-card">

            <span>
              Active
            </span>

            <strong>
              {activeTodos}
            </strong>

          </div>

          <div className="stat-card">

            <span>
              Completed
            </span>

            <strong>
              {completedTodos}
            </strong>

          </div>

        </div>

        {/* =========================
            Todo Section
            ========================= */}

        <section className="todos-section">

          <div className="section-header">

            <h2>
              My Tasks
            </h2>

            <span>
              {filteredTodos.length}{" "}
              {filteredTodos.length === 1
                ? "task"
                : "tasks"}{" "}
              shown
            </span>

          </div>

          {/* =========================
              Search + Filters
              ========================= */}

          <div className="todo-controls">

            <input
              type="text"
              value={search}
              onChange={(e) =>
                setSearch(
                  e.target.value
                )
              }
              placeholder="Search todos..."
              className="search-input"
            />

            <div className="filters">

              <button
                className={
                  filter === "all"
                    ? "filter-active"
                    : ""
                }
                onClick={() =>
                  setFilter("all")
                }
              >
                All
              </button>

              <button
                className={
                  filter === "active"
                    ? "filter-active"
                    : ""
                }
                onClick={() =>
                  setFilter(
                    "active"
                  )
                }
              >
                Active
              </button>

              <button
                className={
                  filter ===
                  "completed"
                    ? "filter-active"
                    : ""
                }
                onClick={() =>
                  setFilter(
                    "completed"
                  )
                }
              >
                Completed
              </button>

            </div>

          </div>

          {/* =========================
              Error
              ========================= */}

          {error && (
            <div className="error">
              {error}
            </div>
          )}

          {/* =========================
              Loading
              ========================= */}

          {loading ? (

            <div className="loading">
              Loading todos...
            </div>

          ) : filteredTodos.length ===
            0 ? (

            /* =========================
               Empty State
               ========================= */

            <div className="empty-state">

              <h3>
                No todos found
              </h3>

              <p>
                {search
                  ? "Try a different search."
                  : filter ===
                    "active"
                  ? "You have no active tasks."
                  : filter ===
                    "completed"
                  ? "You have no completed tasks."
                  : "Create your first task above."}
              </p>

            </div>

          ) : (

            /* =========================
               Todo List
               ========================= */

            <div className="todo-list">

              {filteredTodos.map(
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

      {/* =========================
          Edit Modal
          ========================= */}

      <EditTodoModal
        todo={editingTodo}
        onClose={() =>
          setEditingTodo(null)
        }
        onSave={
          handleSaveEdit
        }
      />

      {/* =========================
          Delete Modal
          ========================= */}

      <DeleteTodoModal
        todo={deletingTodo}
        onClose={() =>
          setDeletingTodo(null)
        }
        onConfirm={
          handleConfirmDelete
        }
        loading={deleteLoading}
      />
        
    </div>
  );
}

export default Dashboard;