import { useEffect, useState } from "react";

function EditTodoModal({
  todo,
  onClose,
  onSave,
}) {
  const [title, setTitle] =
    useState("");

  const [description, setDescription] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  const [error, setError] =
    useState("");

  /*
   * Populate the form when
   * the selected todo changes.
   */

  useEffect(() => {
    if (todo) {
      setTitle(todo.title || "");

      setDescription(
        todo.description || ""
      );

      setError("");
    }
  }, [todo]);

  if (!todo) {
    return null;
  }

  const handleSubmit = async (e) => {
    e.preventDefault();

    setError("");

    if (!title.trim()) {
      setError("Title is required");
      return;
    }

    try {
      setLoading(true);

      await onSave({
        title: title.trim(),
        description:
          description.trim(),
      });

      onClose();

    } catch (error) {
      console.error(
        "Edit todo error:",
        error
      );

      setError(
        error.response?.data?.error ||
          "Failed to update todo"
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="modal-overlay"
      onClick={onClose}
    >
      <div
        className="modal"
        onClick={(e) =>
          e.stopPropagation()
        }
      >
        <div className="modal-header">

          <h2>
            Edit Todo
          </h2>

          <button
            className="modal-close"
            onClick={onClose}
            disabled={loading}
          >
            ×
          </button>

        </div>

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
                setTitle(
                  e.target.value
                )
              }
              disabled={loading}
              autoFocus
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
              rows="5"
              disabled={loading}
            />

          </div>

          <div className="modal-actions">

            <button
              type="button"
              className="secondary-button"
              onClick={onClose}
              disabled={loading}
            >
              Cancel
            </button>

            <button
              type="submit"
              className="primary-button"
              disabled={loading}
            >
              {loading
                ? "Saving..."
                : "Save Changes"}
            </button>

          </div>

        </form>
      </div>
    </div>
  );
}

export default EditTodoModal;