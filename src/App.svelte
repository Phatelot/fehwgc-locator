<script lang="ts">
  import type { Component } from 'svelte';
  import TokenForm from './lib/TokenForm.svelte'
  import Locator from './lib/Locator.svelte';
    import Background from './lib/Background.svelte';

  let hash = $state(getHash());

  function getHash() : string {
    return location.hash.slice(1) || '/';
  }

  // Listen to hash changes
  window.addEventListener('hashchange', () => {
    hash = getHash();
  });

  // Simple navigation helper
  function navigate(path: string) {
    location.hash = path;
  }

  // Route table
  const routes: {[hash: string]: Component} = {
    '/': Locator,
    '/token': TokenForm,
  };

  // Pick the component based on the current hash
  const ComponentToDisplay = $derived(routes[hash] ?? routes['/']);

</script>

<main class="container-fluid">
  <Background/>
  <nav>
    <ul>
      <li>Fire emblem heroes drive</li>
    </ul>
    <ul>
      <li><a href="#/" onclick={() => navigate('/')} class="contrast">Image</a></li>
    </ul>
  </nav>

  <ComponentToDisplay />
</main>

<style>
  li, a {
    color: black;
    -webkit-text-stroke-width: 4px;
  }
</style>
