# Stickian Empire - The Game

![GitHub](https://img.shields.io/github/license/stickianempire/stickian)


*(under construction)*
## Project Description
*(under construction)*
## Project Structure
*(under construction)*

## Project Methodology
In this project, there are some assumptions that one must take into account. Namely, collaborators must be aware not only of the project policies but also of the applied methodology. In the following points, we will be going through the source-control branching model and the commiting terminology.
  
- #### Git Workflow:
  In this project, we adopt the [Trunk Based Development](https://trunkbaseddevelopment.com/) as the source-control branching model. it is is one of a set of capabilities that drive higher software delivery and organizational performance. This model is focused in the development of a single main branch, called "trunk", resisting to any pressure to create other long-lived development branches. In practise, each developer divides their own work into small batches and merges that work into trunk at least once (and potentially several times) a day. Naturally, for a single developer the model is reduced to the direct development of the master branch.
  
  The only exception in this work to this sort of development, is that release branches are not present, assuming the master is the latest and the truest form of the live application.

  **IMPORTANT**: every branch has to be merged through a *PULL REQUEST* and has to be approved by at **least** one developer. Branches are **squashed and merged** into a single meaningful commit.

- #### Branching Nomenclature:
  Following the previously mentioned workflow, in multi-colaborators project, each feature/hot fix branch shall follow the following nomenclature:
  `<identifier>/<work-type>/<description>`

  The `identifier` should be a string that each collaborator should choose and stick to it, in order to ease the development process, namely with auto completion when checking out between branches. The `work-type` should fall between the values of `doc`, `ftr`, `fix`, or any other three letter description of the existing work.

- #### Commits terminology:
  Merge commits must have a single line with the maximum of 50 characters, in imperative form, declaring the changes made, and if possible, their basic purpose. The following message, with each line reaching no more than 70 characters, has to explain the **HOW** and **WHY**, rather than the *what*, since the last question should be answered through code.

  Branch commits do not matter.

*(under construction)*

## Project Policies
(example from [template](https://betterscientificsoftware.github.io/A-Team-Tools/TeamPoliciesTemplate.html))

The following policies are meant to guide collaborators in their activities, establishing expectations for ongoing work.

- #### Code of Conduct:
    1. Collaborators will conduct themselves in a professional manner, following the project's code of conduct, 
       avoiding any non-professional behaviours.
    2. Collaborators will frequently update the project's Kanban boards. Specifically, each collaborator is in charge 
       of managing and updating the tasks to which they are assigned. The tasks priority order must be respected (Hot 
       Fixes, On Progress and To Do), although some ponctual exceptions are allowed, accordingly to the project plan.

*(under construction)*

## Contributors

<table>
  <tbody>
    <tr>
      <td align="center">
        <a href="https://github.com/stickianempire">
          <img src="https://github.com/stickianempire.png" width="100px;" style="border-radius:100%"/>
          <br/><sub><b>Stickian Empire</b></sub><br/>
        </a>
        <a href="mailto: stickianempire@gmail.com" title="Mail">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/mail.png" width="20" style="margin-top:10px"/>
        </a>
      </td>
      <td align="center">
        <a href="https://github.com/bgonilho">
          <img src="https://github.com/bgonilho.png" width="100px;" style="border-radius:100%"/>
          <br /><sub><b>bgonilho</b></sub><br />
        </a>
        <a href="mailto: bernardomiguel0@gmail.com" title="Mail">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/mail.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://www.linkedin.com/in/bernardogonilho/" title="LinkedIn">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/linkedin.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://discord.com/users/318488333233946624" title="Discord">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/discord.png" width="20" style="margin-top:10px"/>
        </a>
      </td>
      <td align="center">
        <a href="https://github.com/luisferreira32">
          <img src="https://github.com/luisferreira32.png" width="100px" style="border-radius:100%"/>
          <br /><sub><b>luisferreira32</b></sub><br />
        </a>
        <a href="mailto: stickianempire@gmail.com" title="Mail">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/mail.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://www.linkedin.com/in/lu%C3%ADs-morgado-ferreira-90a558142/" title="LinkedIn">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/linkedin.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://discord.com/users/279263718486048768" title="Discord">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/discord.png" width="20" style="margin-top:10px"/>
        </a>
      </td>
      <td align="center">
        <a href="https://github.com/MrToino">
          <img src="https://github.com/MrToino.png" width="100px;" style="border-radius:100%"/>
          <br /><sub><b>Mr Toino</b></sub><br />
        </a>
        <a href="mailto: antonio.mf97@gmail.com" title="Mail">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/mail.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://www.linkedin.com/in/ant%C3%B3nio-medeiros-fernandes/" title="LinkedIn">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/linkedin.png" width="20" style="margin-top:10px"/>
        </a>
        <a href="https://discord.com/users/318061313374814219" title="Discord">
        <img src="https://cdn.jsdelivr.net/gh/dmhendricks/signature-social-icons/icons/round-flat-filled/50px/discord.png" width="20" style="margin-top:10px"/>
        </a>
      </td>
    </tr>
  </tbody>
</table>

