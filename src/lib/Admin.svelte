<script lang="ts">
    import { getGithubToken, triggerUpdateImageWorkflow, updateGithubToken } from "./github";

    let token = $state(getGithubToken() || '');
</script>

<h1>Admin</h1>

<button onclick="{() => triggerUpdateImageWorkflow()}">Update picture</button>

<form>
  <label>
    Token
    <input type=password bind:value={token} onchange={() => updateGithubToken(token)}/>
  </label>
</form>


<style>
  label, h1, button {
    color: black;
    -webkit-text-stroke-width: 4px;
  }
</style>
