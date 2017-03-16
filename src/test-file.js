const app =  document.getElementById('app');

const createH1 = (text, id, className) => {
  const h1 = document.createElement('h1');

  h1.id = id;
  h1.classList.add(className);
  h1.innerText = text;

  return h1;
};

const greeting = createH1('Go Server Go!', 'greet', 'intros');

app.append(greeting);
