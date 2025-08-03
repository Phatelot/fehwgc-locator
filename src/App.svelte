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

  function parseRoute(hash: string): { route: string; params: Record<string, string> } {
    const [path, queryString] = hash.split('?');
    const params: Record<string, string> = {};

    if (queryString) {
      const searchParams = new URLSearchParams(queryString);
      for (const [key, value] of searchParams.entries()) {
        params[key] = value;
      }
    }

    return { route: path, params };
  }

  function showAdmin() : boolean {
    return !!getGithubToken();
  }

  // Listen to hash changes
  window.addEventListener('hashchange', () => {
    hash = getHash();
  });

  let routeData = $derived(parseRoute(hash));

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

  {#key routeData.route + JSON.stringify(routeData.params)}
  {#if routeData.route in routes}
    {#if routeData.route === '/'}
      <Locator selectedCharacterSlug={routeData.params["cs"]} />
    {:else if routeData.route === '/admin'}
      <Admin />
    {/if}
  {:else}
    <Locator />
  {/if}
{/key}
</main>

<style>
  li, a {
    color: black;
    -webkit-text-stroke-width: 4px;
  }
</style>
