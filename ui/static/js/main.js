// Получаем элемент select
const selectElement = document.getElementById('banID');

// Добавляем обработчик события change
selectElement.addEventListener('change', (event) => {
// Сохраняем значение выбранного элемента в локальном хранилище
localStorage.setItem('selectedOption', event.target.value);
});
// Получаем сохраненное значение выбранного элемента из локального хранилища
const selectedOption = localStorage.getItem('selectedOption');

// Если сохраненное значение не равно null, устанавливаем его как выбранный элемент
if (selectedOption !== null) {
selectElement.value = selectedOption;
}


// Получаем элемент input
const inputElement = document.getElementById('inputText');

// Добавляем обработчик события input
inputElement.addEventListener('input', (event) => {
// Сохраняем значение поля ввода в локальном хранилище
localStorage.setItem('inputValue', event.target.value);
});
// Получаем сохраненное значение из локального хранилища
const savedValue = localStorage.getItem('inputValue');

// Если сохраненное значение не равно null, устанавливаем его в качестве значения поля ввода
if (savedValue !== null) {
inputElement.value = savedValue;
}