export async function load({ url }) {
  const parent = url.searchParams.get('parent');
  if (parent != null) {
    return { parent: parseInt(parent) };
  } else {
    return { parent: null };
  }
}

