function addTask() {
    const taskInput = document.getElementById('taskInput');
    const taskText = taskInput.value.trim();
    if (taskText === '') return;

    const taskList = document.getElementById('taskList');

    // Create a new list item
    const listItem = document.createElement('li');
    listItem.textContent = taskText;

    // Add a delete button to the list item
    const deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';
    deleteButton.onclick = function() {
        listItem.remove();
        deleteTask(taskText);
    };
    listItem.appendChild(deleteButton);

    // Append the list item to the task list
    taskList.appendChild(listItem);

    // Clear the input field
    taskInput.value = '';

    // Add the task to the backend
    addTaskToBackend(taskText);
}

function addTaskToBackend(taskText) {
    fetch('/tasks/add', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ text: taskText }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to add task');
        }
    })
    .catch(error => console.error('Error adding task:', error));
}

function deleteTask(taskText) {
    fetch('/tasks/delete', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ text: taskText }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to delete task');
        }
    })
    .catch(error => console.error('Error deleting task:', error));
}
