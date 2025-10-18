export function load({ url }) {
  const login = url.searchParams.get('login');
  const signup = url.searchParams.get('signup');
  if (login != null) {
    return { form: "login" }
  } else if (signup != null) {
    return { form: "signup" }
  } else {
    return { form: "login" }
  }
}
