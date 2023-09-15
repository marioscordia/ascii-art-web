
const textarea = document.getElementById('inputText');

function validateTextarea() {
    const text = textarea.value;
    const asciiRegex = /^[\x00-\x7F\n]*$/;

    if (text.trim() === '') {
        textarea.setCustomValidity('Поле не может быть пустым.');
    } else if (!asciiRegex.test(text)) {
        textarea.setCustomValidity('Можно использовать только символы стандартной ASCII таблицы и символ перевода строки.');
    } else {
        textarea.setCustomValidity('');
    }
}

textarea.addEventListener('input', validateTextarea);