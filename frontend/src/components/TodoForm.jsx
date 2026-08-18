import { useState } from "react";

function TodoForm({ onTodoCreated }) {

  const [title, setTitle] =
    useState("");

  const [description, setDescription] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  const [error, setError] =
    useState("");

  const handleSubmit = async (e) => {

    e.preventDefault();

    setError("");

    if (!title.trim()) {
      setError(
        "Title is required"
      );

      return;
    }

    try {

      setLoading(true);

      await onTodoCreated({
        title: title.trim(),
        description:
          description.trim(),
      });

      /*
       * Clear form after successful
       * creation.
       */

      setTitle("");
      setDescription("");

    } catch (error) {

      console.error(
        "Create todo error:",
        error
      );

      setError(
        error.response?.data?.error ||
        "Failed to create todo"
      );

    } finally {

      setLoading(false);

    }
  };

  return (
    <div className="todo-form-card">

      <h2>
        Add New Todo
      </h2>

      {error && (
        <div className="error">
          {error}
        </div>
      )}

      <form
        onSubmit={handleSubmit}
      >

        <div className="form-group">

          <label>
            Title
          </label>

          <input
            type="text"
            value={title}
            onChange={(e) =>
              setTitle(e.target.value)
            }
            placeholder="Enter todo title"
          />

        </div>

        <div className="form-group">

          <label>
            Description
          </label>

          <textarea
            value={description}
            onChange={(e) =>
              setDescription(
                e.target.value
              )
            }
            placeholder="Enter description"
            rows="4"
          />

        </div>

        <button
          type="submit"
          disabled={loading}
        >
          {loading
            ? "Adding..."
            : "Add Todo"}
        </button>

      </form>

    </div>
  );
}

export default TodoForm;