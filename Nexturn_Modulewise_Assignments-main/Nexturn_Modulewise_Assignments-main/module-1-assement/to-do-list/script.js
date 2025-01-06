// Fetch tasks from localStorage or initialize an empty array
let tasks = JSON.parse(localStorage.getItem('tasks')) || [];

function renderTasks() {
    const taskList = document.getElementById('taskList');
    const pendingCount = document.getElementById('pendingCount');
    taskList.innerHTML = '';

    let pendingTasks = 0;
    tasks.forEach((task, index) => {
        const taskItem = document.createElement('li');

        const taskName = document.createElement('span');
        taskName.textContent = task.name;
        if (task.completed) taskName.classList.add('completed');

        taskName.addEventListener('click', () => toggleComplete(index));
        taskItem.appendChild(taskName);

        const taskActions = document.createElement('div');
        taskActions.classList.add('task-actions');

        const editButton = document.createElement('button');
        editButton.textContent = 'Edit';
        editButton.onclick = () => editTask(index);
        taskActions.appendChild(editButton);

        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Delete';
        deleteButton.onclick = () => deleteTask(index);
        taskActions.appendChild(deleteButton);

        taskItem.appendChild(taskActions);

        taskList.appendChild(taskItem);

        if (!task.completed) pendingTasks++;
    });

    pendingCount.textContent = `Pending tasks: ${pendingTasks}`;
    saveTasks();
}

function addTask() {
    const taskInput = document.getElementById('taskInput');
    const taskName = taskInput.value.trim();
    if (taskName) {
        tasks.push({ name: taskName, completed: false });
        taskInput.value = '';
        renderTasks();
    }
}

function editTask(index) {
    const newTaskName = prompt('Edit task:', tasks[index].name);
    if (newTaskName !== null) {
        tasks[index].name = newTaskName.trim();
        renderTasks();
    }
}

function deleteTask(index) {
    tasks.splice(index, 1);
    renderTasks();
}

function toggleComplete(index) {
    tasks[index].completed = !tasks[index].completed;
    renderTasks();
}

function saveTasks() {
    localStorage.setItem('tasks', JSON.stringify(tasks));
}

// Initial render
renderTasks();
