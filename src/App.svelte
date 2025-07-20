<script lang="ts">
  import type { Component } from 'svelte';
  import Admin from './lib/Admin.svelte'
  import Locator from './lib/Locator.svelte';
  import Background from './lib/Background.svelte';
  import { getGithubToken } from './lib/github';

  let hash = $state(getHash());

  function getHash() : string {
    return location.hash.slice(1) || '/';
  }

  function showAdmin() : boolean {
    return !!getGithubToken();
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
    '/admin': Admin,
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
      {#if showAdmin()}
        <li><a href="#/admin" onclick={() => navigate('/admin')} class="contrast">Admin</a></li>
      {/if}
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
