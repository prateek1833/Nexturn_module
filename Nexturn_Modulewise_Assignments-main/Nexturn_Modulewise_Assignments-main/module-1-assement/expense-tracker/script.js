const form = document.getElementById('expense-form');
const expenseTableBody = document.querySelector('#expense-table tbody');
const summaryDiv = document.getElementById('summary');

// To store expenses
let expenses = [];

// Add Expense
form.addEventListener('submit', function (event) {
  event.preventDefault();

  const amount = parseFloat(document.getElementById('amount').value);
  const description = document.getElementById('description').value;
  const category = document.getElementById('category').value;

  if (!amount || !description || !category) return;

  const expense = { amount, description, category };
  expenses.push(expense);
  updateTable();
  updateSummary();

  form.reset();
});

// Update Expense Table
function updateTable() {
  expenseTableBody.innerHTML = '';
  expenses.forEach((expense, index) => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${expense.amount}</td>
      <td>${expense.description}</td>
      <td>${expense.category}</td>
      <td><button class="delete-btn" data-index="${index}">Delete</button></td>
    `;
    expenseTableBody.appendChild(row);
  });
}

// Update Expense Summary
function updateSummary() {
  const summary = {};
  expenses.forEach(expense => {
    if (!summary[expense.category]) summary[expense.category] = 0;
    summary[expense.category] += expense.amount;
  });

  summaryDiv.innerHTML = '';
  for (const category in summary) {
    const p = document.createElement('p');
    p.textContent = `${category}: $${summary[category].toFixed(2)}`;
    summaryDiv.appendChild(p);
  }
}

// Delete Expense
expenseTableBody.addEventListener('click', function (event) {
  if (!event.target.classList.contains('delete-btn')) return;

  const index = event.target.getAttribute('data-index');
  expenses.splice(index, 1);
  updateTable();
  updateSummary();
});
