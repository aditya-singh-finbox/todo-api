function TodoItem({
  todo,
  onToggle,
  onEdit,
  onDelete,
}) {
  return (
    <div
      className={`todo-item ${
        todo.completed
          ? "completed"
          : ""
      }`}
    >

      <div className="todo-main">

        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() =>
            onToggle(todo)
          }
        />

        <div className="todo-content">

          <h3>
            {todo.title}
          </h3>

          {todo.description && (
            <p>
              {todo.description}
            </p>
          )}

          <small>
            {todo.completed
              ? "Completed"
              : "Pending"}
          </small>

        </div>

      </div>

      <div className="todo-actions">

        <button
          onClick={() =>
            onEdit(todo)
          }
        >
          Edit
        </button>

        <button
          className="delete-button"
          onClick={() =>
            onDelete(todo.id)
          }
        >
          Delete
        </button>

      </div>

    </div>
  );
}

export default TodoItem;