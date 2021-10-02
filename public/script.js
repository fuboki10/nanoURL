const backendUrl = `${window.location}urls/`;
const errorMessage = "Something wrong happended <br/>"
const successMessage = "CONGRATULATIONS, here is your new link <br\>"

const handleError = () => {
  const span = document.getElementById('output');
  span.innerHTML = errorMessage;
}

const handleSuccessfullRequest = (url) => {
  const span = document.getElementById('output');
  const link = `<a href="${url}" target="_blank">${window.location + url}</a>`;
  span.innerHTML = successMessage + link;
}

const sendCreateUrlRequest = async (long_url) => {
  const response = await fetch(backendUrl, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json'
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify({ long_url })
  });

  return response.json();
}

const btnClickHandler = async () => {
  const url = document.getElementById('url');
  try {
    const shortUrl = (await sendCreateUrlRequest(url.value)).url;
    handleSuccessfullRequest(shortUrl);
  } catch (error) {
    console.log(error);
    handleError();
  }
}

const main = () => {
  const btn = document.getElementById('shorten-btn');
  btn.addEventListener('click', btnClickHandler)
}







document.addEventListener('DOMContentLoaded', main);