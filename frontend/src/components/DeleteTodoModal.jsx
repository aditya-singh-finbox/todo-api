function DeleteTodoModal({
  todo,
  onClose,
  onConfirm,
  loading,
}) {
  if (!todo) {
    return null;
  }

  return (
    <div
      className="modal-overlay"
      onClick={onClose}
    >
      <div
        className="modal delete-modal"
        onClick={(e) =>
          e.stopPropagation()
        }
      >

        <div className="modal-header">

          <h2>
            Delete Todo
          </h2>

          <button
            className="modal-close"
            onClick={onClose}
            disabled={loading}
          >
            ×
          </button>

        </div>

        <div className="delete-content">

          <p>
            Are you sure you want to
            delete this todo?
          </p>

          <strong>
            {todo.title}
          </strong>

          <p className="warning-text">
            This action cannot be
            undone.
          </p>

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
            type="button"
            className="delete-confirm-button"
            onClick={onConfirm}
            disabled={loading}
          >
            {loading
              ? "Deleting..."
              : "Delete"}
          </button>

        </div>

      </div>
    </div>
  );
}

export default DeleteTodoModal;