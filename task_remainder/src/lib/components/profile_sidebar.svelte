<script lang="ts">
  const params = {
    name: "Andrey Carvalho",
    description: "Lorem ipsum dolorLorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor Lorem ipsum dolor ",
    picture: "https://raw.githubusercontent.com/ancarvalho/curriculum-vitae/master/static/profile.jpg",
    links: {
      github: "https://github.com",
      LinkedIn: "https://linkedin.com",
      email: "mailto:ancarvalho6@gmail.com",
      twitter: "https://twitter.com",
    },
  };

</script>

<div class="flex flex-col justify-center items-center gap-4">
  
  <img class="rounded-full h-32" src={params.picture} alt="" />
  <div >
    <p class="text-xl text-center pb-4">
      {params.name}
    </p>
    <p class="text-sm text-center line-clamp-6">
      {params.description}
    </p>
  </div>
  <ul class="uppercase flex flex-col items-center gap-2">
    {#each Object.entries(params.links) as link}
      <li>
        <a href={link[1]}>
          {link[0]}
        </a>
      </li>
    {/each}
  </ul>
</div>


