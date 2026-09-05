# Potatoverse — mission, vision, and product decisions

Last updated: 2026-09-06

This is the project's living planning document. Read it before planning or implementing product behavior. It records the agreed concept, a proposed complete user journey, and questions the project owner needs to answer.

## How to maintain this document

- **Confirmed** means the owner has explicitly agreed to it. **Proposed** means it is a working suggestion, not an approved requirement.
- Do not silently turn a suggestion into a product decision. A prototype can explore a proposal if its assumptions are clearly identified.
- Each unanswered question below includes a suggestion and its reason. The owner can answer using the question ID, for example, `Q01: Allow membership in multiple societies`.
- After an answer arrives, record the decision in the confirmed decision log, retain its question ID for traceability, and **remove that question and its suggestion from the unanswered section**. Do not preserve the suggestion elsewhere as if it were a requirement.
- Update the affected flows, permissions, and development tasks in the same edit. If an answer leaves a separate issue unresolved, create a focused follow-up question with a suggestion.
- If an answer replaces an earlier decision, update the current decision and mark the previous decision superseded. Ask about actual conflicts rather than silently choosing between them.
- Do not describe suggested interfaces, providers, policies, election rules, or visual styles as already decided.
- Develop one useful part at a time. Resolve questions needed for the current part; future elections, governance types, or competitions must not block unrelated work. Keep the broader map for future planning.

## Mission and vision

### Mission

Build one social media world where people create and join societies, connect around shared interests, choose how their societies are governed, and represent them in challenges and competition.

### Vision

Potatoverse is **one connected social app containing societies created by its users**. Each society has its own identity and a governance type. **Kingdom** and **Republic** are the initial governance types: they define how authority is selected, distributed, and used for moderation. More types may be introduced as the product develops.

A country, interest, or group of people can be associated with multiple societies using different governance types. For example, UK-based members could create one Kingdom society and another Republic society. A society's identity and its governance type are separate choices.

People discover, create, and join societies within the shared app. Societies develop their own communities and leaders, interact with others, and can challenge each other in esports, other online activities, or real-life events. The competition formats are deliberately undecided.

The exact creation rules, content policies, powers, membership boundaries, and competition mechanics remain open. Build the experience incrementally: each completed part should be usable or reviewable before proceeding to the next.

### Confirmed decision log

| ID | Confirmed decision |
| --- | --- |
| C01 | The project is a social media app called Potatoverse. |
| C02 | **Superseded by C10:** initial onboarding was previously framed as choosing between two fixed societies. Onboarding now needs discovery/joining and society creation. |
| C03 | Kingdom and Republic define different governance and content moderation systems. Their specific content rules are not yet defined. |
| C04 | In a society using the Kingdom type, people appoint its king. That king assigns moderators to the society's topics, groups, or channels. The appointment process is not yet defined. |
| C05 | In a society using the Republic type, people elect representatives who lead its communities. Representatives can progress through further votes to that society's presidency. The number of levels and electorate at each level are not yet defined. |
| C06 | The basic social unit—topic, group, channel, or community—is still undecided. |
| C07 | This root-level `agent.md` stores the mission, vision, proposed flows, open questions, and owner decisions. Suggestions are removed when their questions are answered. |
| C08 | Societies share one connected app and can challenge one another and compete. The earlier fixed Kingdom-versus-Republic pairing is superseded by C10. |
| C09 | Competitions may include esports, other online activities, or real-life events. No specific format has been selected yet. |
| C10 | Users can create societies. Kingdom and Republic are governance types that a society can use; they are not the only two society identities in Potatoverse. |
| C11 | Multiple societies can serve people from the same country or background while choosing different governance types. |
| C12 | The direction allows for additional governance types in the future; their mechanics and timing remain open. |
| C13 / Q10 (delivery approach) | Develop incrementally, one part at a time. A complete implementation of every governance type and competition feature is not a prerequisite for the first development checkpoint. Q10 now asks only about that checkpoint's scope. |

### Core distinction — confirmed model, illustrative examples

| Example society name | Governance type | Example audience |
| --- | --- | --- |
| Crown & Chips | Kingdom | UK-based members who prefer delegated royal leadership. |
| Common Ground UK | Republic | UK-based members who prefer elected representation. |
| Night Owl Gamers | Kingdom | A gaming interest group. |

These names illustrate the model; they are not seeded societies or approved branding. Leaders and elections belong to a particular society. No single king or president governs every society using that type.

## Working vocabulary — proposed

These terms make the draft readable; they do not settle the naming questions.

| Term | Working meaning |
| --- | --- |
| Account | A person's platform identity and sign-in credentials. |
| Society | A user-created social organization with its own name, identity, members, rules, and selected governance type. |
| Governance type | A supported governing model, initially Kingdom or Republic, that defines offices, leadership selection, and authority. Many societies can use the same type. |
| Founder | The person who creates a society. Founding setup permissions and the first governing office are separate decisions. |
| Society membership | An account's affiliation, governance rights, and potential competition representation. This is separate from permission to interact socially across society boundaries. |
| Community | A space for a shared interest within one society. Used provisionally instead of topic/group/channel. |
| Member | Someone participating in a society or community; political participation need not be mandatory. |
| King | A particular Kingdom-type society's ruler, selected by its people through the process still to be decided. |
| Representative | An elected community leader within a particular Republic-type society. |
| Higher office | A possible elected position between representative and president; the existence and number of these positions are open. |
| President | The highest elected position within a particular Republic-type society. Its powers are open. |
| Moderator | A person with specific content-handling powers over a defined space. Republic staffing is still open. |
| Platform operator | The team maintaining the service and any approved platform-wide rules. This is distinct from a user-held political office. |
| Shared space | A platform-wide area, such as a challenge page, where members of different societies interact under an explicitly stated moderation policy. |
| Challenge | A competition proposal linking specific participating societies, with agreed rules and an acceptance process. |
| Competition/event | An accepted challenge with participants, a schedule, results, and a dispute process; its format may be online or offline. |

## Perspectives the product must account for — proposed

| Perspective | What the experience needs to resolve |
| --- | --- |
| New member | Understand the practical difference between societies before committing, then find people and content quickly. |
| Society founder | Choose a name and supported governance type, establish initial rules, invite members, and understand how founding authority ends or transfers. |
| Casual member | Enjoy the social app without campaigning, running for office, or voting. |
| Politically active member | Know who can vote, what an office controls, how to stand for office, and why an election result is valid. |
| Losing candidate or minority group | Continue participating after losing and have a route to challenge misconduct by the majority or leader. |
| Reporter and accused member | Receive a clear process, reasons for outcomes, appropriate privacy, and a way to challenge mistakes. |
| Moderator or community leader | Understand their jurisdiction, workload, authority, term, and handover responsibilities. |
| King or president | Have meaningful, bounded powers and an explicit process for succession, inactivity, and removal. |
| Competitor or spectator | Represent or support a society, understand the rules, find shared events, and see results without having to hold office. |
| Organizer or referee | Obtain both sides' agreement, fill teams fairly, manage cancellations, and resolve results without one society controlling its opponent. |
| Platform operator | Support many independently governed societies, handle service failures and creation abuse, and keep sensitive account data separate from political authority. |
| Returning, departing, or mobile member | Resume interrupted setup, understand changed rules, leave predictably, and use accessible controls on a small screen. |

## End-to-end user flows — proposed

This maps the intended experience, including failure and recovery paths. It is ready for discussion, not a claim that all requirements are settled. Question IDs identify the decisions that must resolve each flow.

### F01 — First visit, account creation, and society discovery

1. Arrive at a welcome page that explains the shared social world, user-created societies, governance types, and cross-society challenges.
2. Sign up or sign in. Preserve the destination when arriving through a community invitation or shared post.
3. Complete required verification and a minimal profile. Recover from an existing email, expired link, failed provider sign-in, or interrupted session.
4. Discover named societies through search, interests, optional region/language filters, and governance-type filters. Inspect a society's actual rules, leaders/founding status, and membership conditions. A governance-type label alone does not identify a society.
5. Join a chosen society and acknowledge its rules, or create a society through F11. Apply the approved membership model and explain any effects on existing affiliations. Do not interpret an accidental card tap as final enrollment.
6. Pick interests and join suggested communities, with a way to skip recommendations.
7. Arrive at the shared app home with a clear next action from the features currently available: discover a society, join a conversation, make a first post, or view a challenge.

Persist onboarding progress. A returning member goes to their home; an incompletely onboarded member resumes the missing step. An invitation identifies a specific society and needs an explanation consistent with the membership policy, rather than silently enrolling the person. Show an honest empty directory and a creation path when no societies match. Whether joining can be skipped is part of Q01.

Dependencies for the entry flow: Q01, Q10, Q11, Q12, Q26, Q36, Q38, Q39. Add Q02/Q13/Q30 when the feed and communities are implemented; competition rules do not block account or discovery work.

### F02 — Everyday social participation

1. Open one shared home feed, with clear society/community labels and proposed filters for followed interests, one's own society, and activity across Potatoverse.
2. Discover people, communities, and challenges across the app; inspect a space's rules, leadership, and participation conditions before joining its conversation.
3. Read posts, open a discussion, create a post, comment, and react using the formats approved for the first release.
4. Receive relevant replies and moderation notices. Discover cross-society events and access governance activity while keeping ordinary conversation easy to reach.
5. Edit or delete eligible own content; mute or block people and report content when needed.

Handle empty feeds, no search results, failed submissions, duplicate retries, removed posts, and lost permissions. If rules or permissions change while someone is writing, explain the change and preserve their draft where possible. Visibility and blocking behavior must be consistent across feed, search, profiles, and notifications.

Dependencies: Q02, Q10, Q13, Q15, Q16, Q25, Q30.

### F03 — Community lifecycle

1. Find an existing community or request/create one under the society's creation policy.
2. Define its name, purpose, society, rules, and initial leadership status.
3. In Kingdom, obtain a moderator appointment through the king's process. In Republic, start the approved representative-selection process.
4. Open membership and posting when minimum operating conditions are satisfied.
5. Show current leaders, relevant powers, applicable rules, and upcoming selection events on the community page.
6. Handle renaming, inactivity, leader replacement, and archiving while preserving access to relevant history.

Creating a community must not accidentally grant permanent authority that bypasses its society's governance. Duplicate names across societies need clear identity and navigation. An archived community needs defined behavior for posts, reports, and any outstanding leadership contest.

Dependencies: Q02, Q07, Q14, Q21, Q22, Q24.

### F04 — Kingdom leadership and moderator appointment

1. A member opens a specific Kingdom-type society and views its current king, selection history, powers, and any term or review date.
2. When selection opens, eligible people participate in the agreed nomination and appointment process.
3. Finalize the selection, handle any result challenge, and activate the new king's permissions at a defined time.
4. The king chooses a community and nominates an eligible moderator. The nominee accepts before receiving powers.
5. The moderator receives permissions limited to the approved scope and can access the relevant moderation queue.
6. Appointments, resignations, revocations, and succession follow the agreed rules and produce an appropriate record.

Cover declined or expired invitations, self-appointment, vacant communities, unavailable rulers, ruler removal, and whether a new king inherits existing moderators. Members need a route to report abuse by the king or their appointees.

Dependencies: Q04, Q05, Q07, Q08, Q17, Q21, Q22, Q24.

### F05 — Republic representation and advancement

1. A member opens a community in a specific Republic-type society and sees its representative, powers, term, and next election.
2. Eligible candidates nominate themselves or accept a nomination; members can review comparable candidate statements.
3. Eligible voters cast a ballot during the published window and receive confirmation without exposing a private choice.
4. Close voting, resolve ties or disputes, publish the permitted result information, and activate the winning representative's permissions.
5. Representatives who meet the next office's requirements can enter a further election. Advancement requires another vote; it is not an automatic promotion.
6. Repeat for any approved intermediate level within this society. The final election installs that society's president with defined powers.
7. When a leader advances, loses, resigns, or completes a term, transfer authority and fill any resulting vacancy.

The full ladder and electorate at each stage remain open. Do not implement a council or let the president appoint community representatives unless the owner explicitly chooses those rules. Cover uncontested races, zero candidates, low turnout, ties, withdrawn candidates, vote manipulation reports, delayed elections, and a winner who declines office.

Dependencies: Q04, Q06, Q07, Q09, Q18, Q19, Q20, Q21, Q22, Q23, Q24.

### F06 — Rules, moderation, and appeals

1. Members can find the applicable platform, society, community, or shared-event rules before posting, including which rule takes precedence. Cross-society visitors see the same rules as other participants in that space.
2. A member reports content or behavior, selects a reason, and receives a receipt. Blocking or muting can provide immediate personal control.
3. Route the report to an authorized reviewer. Protect reporter identity and private evidence according to the approved policy.
4. The reviewer assesses the relevant rule and evidence, then dismisses, warns, removes content, or restricts participation within their authority.
5. Notify affected users with a reason, scope, duration, and available appeal route. Tell the reporter the permitted outcome information.
6. Route an appeal to a reviewer who was not responsible for the original decision. Record an upheld or reversed result and restore access/content where applicable.
7. Escalate conflicts of interest, leader misconduct, or platform-level cases through the explicitly approved authority structure.

Cover reports against moderators and rulers, unavailable reviewers, repeat reports, edited/deleted evidence, expired restrictions, and rules changed after a post was published. A content suspension's effect on voting must be explicit so leaders cannot silently redefine the electorate through moderation.

Dependencies: Q03, Q04, Q08, Q09, Q15, Q16, Q17, Q22, Q25, Q30, Q33.

### F07 — Rule changes and public accountability

1. Members view current rules and a readable history of changes.
2. An authorized person proposes or issues a change through their society's agreed process.
3. Complete any required review or vote, announce the change, and state when it takes effect.
4. Record who authorized it and the relevant result; apply the correct rule version during moderation.
5. Publish leadership and governance records at an appropriate level of detail, while keeping ballots and private case evidence protected.

The difference between governance types must be visible in these procedures and permissions. Do not infer that a society has unrestricted content or that its king/president has access to private account data.

Dependencies: Q03, Q04, Q08, Q15, Q17, Q18, Q25.

### F08 — Returning, switching, leaving, and account recovery

1. Sign in and return to the last permitted destination without repeating society selection.
2. Recover an account through the chosen authentication method and regain access without duplicating membership or votes.
3. Review profile, notification preferences, memberships, eligibility, current offices, and active restrictions.
4. Before leaving or switching a society, see the effects on content, offices, pending votes, memberships, competition rosters, and any waiting period.
5. Complete the change; resign or transfer offices through the agreed process and show the new participation state.
6. Offer account deletion with a clear explanation of what happens to public content and records needed for existing governance cases.

Enforce the same restrictions after recovery or re-entry. Leaving must not silently erase an election record, create a second ballot, or end an unresolved moderation case. These records' retention and visibility are decisions, not assumptions.

Dependencies: Q01, Q09, Q11, Q12, Q21, Q23, Q25, Q34.

### F09 — Founding and ongoing operation

1. Each newly created society starts under its approved founding policy, with clear rules, available communities when enabled, and visibly identified interim leadership if needed.
2. Explain how and when interim authority ends and ordinary leadership selection begins.
3. Schedule elections and term transitions with published dates and time zones; show eligibility and upcoming actions to members.
4. Handle vacancies, unprocessed reports, failed notifications, and service outages without secretly extending powers or changing votes.
5. Make interventions attributable and reviewable at the appropriate privacy level.

The first member cannot elect a functioning government alone without a founding policy. Every new society needs that transition, rather than relying on a one-time platform launch. Until the policy is approved, development can use clearly labeled demonstration accounts and elections. During incremental development, society drafts can be reviewed before a governance type is available for live operation.

Dependencies: Q07, Q08, Q19, Q21, Q22, Q28, Q36, Q37, Q39.

### F10 — Cross-society challenge, participation, and result

1. Discover a shared challenge hub showing proposals, upcoming events, active competitions, and past results across participating societies.
2. An eligible member or official chooses the proposing society and an opponent, then specifies the activity, team/participant limits, schedule, judging method, conduct rules, and any reward. Clearly distinguish an informal proposal from an official society commitment.
3. The authorized people on each side negotiate and accept the same version of the terms. A rejected or expired proposal remains understandable to its author; material edits require renewed acceptance.
4. Recruit participants under published eligibility and team-selection rules. Confirm which society each person represents and any roster lock before the event.
5. Run the event online, through an external game/activity, or at a real-life location if that format is approved. Support spectators and event discussion in a shared space.
6. Submit results with the required evidence. A designated reviewer or both sides confirm the outcome; challenged results remain provisional.
7. Resolve disputes through the agreed neutral process, finalize the result, and publish shared event history. Award any approved recognition only after finalization.
8. Let participants discuss the event and propose a rematch without requiring a change of society or account.

Cover one society having far more members, insufficient participants, no-shows, cheating accusations, roster changes, affiliation switches, leader turnover, cancelled events, missing evidence, and a judge with a conflict of interest. A competition win does not automatically grant moderation powers or political office. Any connection between results and governance would need a separate owner decision.

Keep competition result disputes separate from reports about comments or harassment around the event. Neither society's leaders should receive unilateral control over shared results or their opponent's members. Offline logistics and esports integrations remain format-specific future work until selected.

Dependencies: Q01, Q04, Q08, Q12, Q16, Q17, Q25, Q30–Q35. A challenge connects society identities; whether their types match is governed by Q31. It need not wait for every governance type to be implemented.

### F11 — Create, establish, and maintain a society

1. From discovery or the shared navigation, choose “Create society.” Check the approved creation eligibility and any existing membership/creation limits before starting.
2. Enter a society name, description, and optional image, interests, region, and language. Validate its unique address and explain that a similar display name may already exist.
3. Choose an available governance type. Explain its offices, founding process, powers, and any limits on changing type later. Types still under development may be described but cannot create a live society with nonfunctional governance.
4. Set the allowed founding rules and membership options using the supported type's controls. Review a summary of what prospective members will see.
5. Save a resumable draft or establish the society under its approved founding policy. The founder does not silently become a permanent king or president.
6. Share an invitation, welcome members, and complete the first leadership-selection process when its conditions are met. Add the first community when that feature is available.
7. Maintain the society's identity, rules, and leadership through authorized changes. Handle founder departure, inactivity, and archiving with explicit effects on memberships, offices, content, elections, and accepted challenges.

Creating a second society of an existing type must work without changing application code or inheriting the first society's leaders. Country and language are discovery information, not a rule that assigns a country's users to one society. Cover abandoned drafts, repeated submissions, duplicate addresses, impersonating names, creation spam, missing eligible leaders, and any limits on changing the governance type.

Dependencies: Q01, Q03, Q04, Q07, Q08, Q10, Q12, Q25, Q36–Q39. Only dependencies for the currently implemented step need to be resolved before that step is built.

## State and permission model — proposed planning constraints

Keep identity, membership, office, and moderation restrictions separate. A single account-level `role` cannot express a person who moderates one community, is a normal member elsewhere, and later leaves office.

| Object | States/transitions to design |
| --- | --- |
| Account/onboarding | Verification pending → profile/setup incomplete → ready; recovery, restriction, and deletion paths. |
| Governance type | Planned → available for new societies → deprecated if needed; versioned behavior for existing societies. A future type is not selectable for live creation until supported. |
| Society | Draft → founding → active → archived; suspension is a separately recorded platform action. Type conversion, if supported later, is an explicit migration. |
| Society/community membership | Not joined → pending if required → active → left or restricted; office and voting eligibility evaluated separately. |
| Community | Proposed → founding → active → archived; leadership may be vacant without erasing the community. |
| Appointment | Offered → accepted/declined/expired; an accepted appointment creates a scoped office or moderation assignment. |
| Election | Scheduled → nominations → voting → counting/review → certified; explicit postponed, cancelled, and rerun paths. |
| Office/term | Vacant → active → expired/resigned/removed; scheduled successor and temporary caretaker rules as needed. |
| Post/comment | Draft → published → edited, author-deleted, or moderator-removed; restoration subject to the approved policy. |
| Moderation case | Submitted → triaged → decided → appealed if eligible → resolved; restriction expiry tracked independently. |
| Rule version | Draft/proposed → approved → effective → superseded. |
| Challenge | Draft → proposed → mutually accepted; rejected, expired, or withdrawn alternatives; changed terms require renewed acceptance. |
| Competition | Registration → scheduled/roster locked → active → result submitted → review/dispute → finalized; postponed or cancelled alternatives. |
| Competition entry | Applied/invited → accepted/waitlisted → roster locked → completed/withdrawn/disqualified, with society represented recorded for that event. |

Every privileged action needs an actor, an explicit scope, an active permission, and the applicable rule. Enforce permissions on the server, including office expiry and revocation. UI visibility alone does not grant authority.

Likely domain objects are accounts, profiles, societies, governance type definitions/versions, memberships, communities, posts, comments, reactions, rule versions, offices, assignments, elections, candidacies, ballots, reports, moderation actions, appeals, challenges, accepted event terms, competitions, teams/entries, results, result disputes, notifications, and audit events. These are planning concepts, not an approved database schema or a requirement to build every object at once.

Use one account and social graph across the product. Attach society and moderation scope to the relevant memberships and spaces. Challenge/event objects link the participating society IDs. Keep voting eligibility, social participation, and competition eligibility distinct.

Store a society's identity separately from its governance type, for example `society.id` and `society.governance_type`. Scope communities, offices, elections, memberships, and moderation actions by society identity. Two Kingdom societies have different kings and electorates. A check such as “user is a king” is insufficient without checking which society and whether the relevant office is active.

Implement supported governance types through a small, explicit set of policies and permissions. Defer a general-purpose constitution editor or user-written governance code. Adding a later type should not require redefining society identity or rebuilding ordinary social features.

## Unanswered questions and suggestions

For the next development checkpoint, prioritize **Q01, Q10, Q36, and Q39**: membership, the first small deliverable, society creation conditions, and the first governance type. Resolve account/provider details when connecting real sign-in. Other questions belong to the feature that needs them; the full list is not a prerequisite for starting development. Every recommendation here is provisional.

References to a king, president, representative, or electorate below always mean those of a specific society using the relevant type. They do not describe platform-wide offices or one shared electorate for every society of that type.

### Foundation decisions

**Q01 — Can a person join multiple societies, and may they use the app before joining any?**

- **Suggestion:** Allow one account to join multiple societies and browse before joining. Give each society its own membership and eligibility rules; joining one must not silently join every society of that type. Use an active society for navigation, not as a limit on all memberships. For competitions, require choosing one represented society per event. Decide whether multi-society members may vote or hold office in several societies when those features are built.

**Q02 — What is the basic social space, and how is it organized?**

- **Suggestion:** Call it a **community**, with posts and comments inside it. Each community belongs to a named society, and the same interest can appear in many societies of any governance type. Start with platform → society → community → discussion. Defer nested groups and live chat channels; the first society-creation checkpoint can end at a minimal society home before communities are built.

**Q03 — What should actually differ between Kingdom and Republic moderation?**

- **Suggestion:** Make authority and rule-making different first: a Kingdom-type society uses its king and appointed moderators; a Republic-type society uses elected leadership. Offer starter content rules that each society can adapt within the approved platform boundaries. Governance type determines the procedure for governing, while the society's published rules explain what content is allowed. Build the procedure for the first selected type before adding the next.

**Q04 — What can each office do, and which powers are reserved to the platform?**

- **Suggestion:** Approve a capability matrix before implementation. Give the king society rule-setting and community moderator appointments. Give representatives community moderation and local rule proposals. Give the president society coordination and the agreed rule-change powers. Define official challenge approval separately in Q31. Reserve account access, private account data, platform-wide enforcement, and emergency service controls to operators. Never grant a president the king's appointment powers by default.

**Q05 — How do people appoint the king?**

- **Suggestion:** Let eligible Kingdom members nominate candidates and vote to select a king for a defined term. The king then appoints moderators rather than having each community elect them. This preserves the delegated hierarchy while making the ruler's initial mandate understandable. Inheritance, indefinite tenure, or a council appointment would require different succession flows.

**Q06 — What is the Republic's exact election ladder, and who votes at each level?**

- **Suggestion:** Start with two elected levels: community members elect representatives; all eligible Republic members elect a president from eligible representatives. This provides advancement through a second election with a manageable pilot. If an intermediate council is essential, define its seats, constituencies, voters, and powers now. Also choose whether a person may vote for representatives in multiple communities; start with one declared home community for electoral representation while allowing social participation in others.

**Q07 — What authority does a founder have, and how does each new society reach its first legitimate leadership?**

- **Suggestion:** Give the founder clearly labeled, temporary setup powers to prepare the society and invite members. Do not automatically award a permanent kingship or presidency. Publish the initial selection process, minimum participation, founding deadline, and fallback if too few members join. Complete the first type's founding policy before activating real governing powers; until then, a saved society can remain a draft. Establishing a society is distinct from creating a community inside it under Q14.

**Q08 — Are there platform-wide rules above all societies and governance types, and who enforces them?**

- **Suggestion:** Yes. Define a small, explicit platform rule set and a separate operator review route. Societies can add rules within that boundary. Explain the rule hierarchy to users and record operator interventions. The actual prohibited content and enforcement policy need owner approval before accepting real public content.

**Q09 — Who may vote, and how do we limit duplicate accounts or moderation being used to suppress voters?**

- **Suggestion:** For the invite-only pilot, require a verified account, a valid society membership, and joining the relevant electorate before nominations open. Freeze the eligible voter list for each election and enforce one ballot per eligible account per contest. Do not let community moderators unilaterally remove someone from that list. Treat challenged eligibility as a separately reviewable case. Account verification alone does not establish one human per account; public-launch eligibility needs a stronger, explicitly chosen approach.

**Q10 — What should our first small development checkpoint contain?**

- **Suggestion:** First build a reviewable entry flow: login/sign-up screens → society discovery → society details → join or create → minimal society home. Society creation captures a name, description, and governance type. Begin as a clearly labeled local prototype if real authentication and founding policies are still undecided, then connect persistence and permissions as the next small task. Communities, elections, and challenges can follow independently. Incremental delivery is confirmed in C13; the contents of this first checkpoint remain open.

### Accounts and social experience

**Q11 — How should people sign in, and can they use pseudonyms?**

- **Suggestion:** Allow public pseudonyms and keep email private. Use a managed authentication provider with email/password, email verification, and account recovery for the pilot; add Google sign-in if desired. Choose the provider in Q28. Use stronger authentication for office holders before they gain sensitive powers.

**Q12 — What happens when someone switches or leaves a society?**

- **Suggestion:** Distinguish changing the society currently viewed from ending a membership. Let a member leave a particular society after explicitly resigning active offices there; retain any other memberships. Keep old content in its original space under its rules, and preserve ballots/case records under the approved privacy policy. Apply election and competition eligibility rules when those features exist. Leaving a society must not delete the platform account. If Q01 chooses a single-membership limit, add an explicit leave-and-join flow.

**Q13 — What content, feed, discovery, and interactions should launch first?**

- **Suggestion:** Begin with text posts, comments, and one simple reaction. Use one chronological feed of joined communities plus shared challenge updates, a platform-wide directory/search, and optional society filters. Allow cross-society discovery and interaction under Q30, independently from governance eligibility. Defer media uploads and person-following until the basic community loop works.

**Q14 — Who can create, join, or archive communities, and are they public or private?**

- **Suggestion:** Use public-to-members communities with open joining during the pilot; operators seed the first few and accept creation requests. Founders receive no permanent governing privilege. Archive rather than erase inactive spaces, with clear notice and preserved case history. Decide public web visibility separately before letting unauthenticated visitors read posts. Defer private communities and approval-based joining.

### Moderation and accountability

**Q15 — Who writes and changes society/community rules, and when do changes apply?**

- **Suggestion:** Let the king approve Kingdom society rules; let Republic representatives approve society rules by a recorded majority vote, with the president publishing the result. Community leaders can propose compatible local additions. Version rules and announce an effective date; normally assess content using the version in force when it was posted. Handle urgent platform interventions through Q08. The Republic procedure needs quorum and tie rules before implementation.

**Q16 — What moderation actions exist, and how are reports handled?**

- **Suggestion:** Start with dismissing a report, issuing a warning, removing content, and time-limited community posting restrictions. Require a rule reference and reason. Keep reporter identities private from the reported person and public viewers. Provide case status to the reporter and action details to the affected member. Reserve society/account-wide restrictions to explicitly authorized reviewers; a moderator's community powers should not silently affect unrelated spaces.

**Q17 — Who hears appeals, including appeals against the king, president, or platform operator?**

- **Suggestion:** Use an operator-managed appeals queue during the pilot, with a reviewer different from the original decision-maker. Apply the society's approved rules unless a platform rule takes precedence. Record conflicts of interest and an escalation route for complaints about operators. Do not promise independent appeal review with only one available operator; settle staffing and the final authority before launch.

### Elections and leadership lifecycle

**Q18 — How are ballots counted, and which election details are public?**

- **Suggestion:** For single-seat pilot elections, use one candidate choice per voter and highest vote count wins, subject to Q19. Keep individual choices private from members and office holders; publish candidates, timetable, total turnout, aggregate results, and certification status. A receipt confirms participation without revealing the choice. Allow replacement of a ballot until closing, with only the final choice counted. This does not promise cryptographic anonymity from platform operators.

**Q19 — What are the election schedule, turnout requirement, tie rule, and dispute process?**

- **Suggestion:** As a pilot starting point, use three days for nominations, three days for voting, and a 24-hour challenge window before office activation. Use a runoff for tied leaders. A race with no candidates remains vacant under a limited caretaker; an uncontested candidate still faces a confirmation vote. Publish a minimum turnout before opening nominations and postpone an election that misses it. Set its numeric threshold against the recruited pilot electorate, then document runoff, repeated-failure, and outage rules before live elections.

**Q20 — Who can stand for office, and how should campaigning work?**

- **Suggestion:** Require eligible membership, an account in good standing under a reviewable policy, candidate consent, and one active candidacy at a time. Provide every candidate the same profile/statement space and an accessible election discussion. For president, require representative status as defined in Q06 and Q23. Start without paid promotion or purchasable voting influence.

**Q21 — How long are terms, can leaders be re-elected, and who acts during vacancies?**

- **Suggestion:** Pilot renewable 90-day king and president terms and 30-day representative terms. These are starting proposals, not evidence-based targets. Use limited, time-bounded caretakers and scheduled replacement selection for vacancies. Expire old permissions automatically and define the handover time. Preserve existing moderator assignments during a royal transition until reviewed, so communities do not abruptly lose coverage.

**Q22 — How can members remove or challenge an abusive or inactive leader?**

- **Suggestion:** Provide a documented recall process initiated by a member petition, followed by a vote of the relevant electorate. Use the same eligibility controls as ordinary elections. Allow temporary operator suspension for platform-rule violations, with independent review. Define petition threshold, recall turnout, cooldown, inactivity period, and caretaker limits before enabling removal; leaders should not decide cases about their own misconduct.

**Q23 — Can one person hold multiple offices within or across societies, and what happens when a representative advances?**

- **Suggestion:** Initially allow one governing office per account across the platform to keep conflicts and workload manageable, while keeping every permission scoped to the actual society. This is separate from ordinary multi-society membership. A representative who becomes their society's president relinquishes their community seat at handover and triggers a replacement election. Fix nomination eligibility for the contest; specify how misconduct removal or resignation affects it. Revisit the cross-society office limit after the pilot.

**Q24 — In Republic, does the elected representative moderate personally or appoint a team? How are moderator powers withdrawn?**

- **Suggestion:** Have representatives moderate their own communities in the first pilot. Add scoped, publicly listed assistants later if workload requires it. Kingdom moderators must accept appointment and can resign or be removed through a recorded process. Revocation must take effect immediately on the server, including for someone with an already-open moderation screen.

### Launch and implementation

**Q25 — What should users control about privacy, notifications, blocking, and deletion?**

- **Suggestion:** Show public handles, community content, office history, and finalized event results; keep email, private ballots, reporter identities, and case evidence out of public profiles. Start with in-app notifications for replies, appointments, elections, challenge invitations/schedule changes, and case outcomes, with preferences for optional reminders. Make blocking hide ordinary interaction and direct targeting, while preserving access to necessary governance/event notices and authorized case review. Before deletion, explain whether content is removed or anonymized and what limited records remain, for how long, and who can access them.

**Q26 — Who is the initial audience, and what age, language, and launch boundaries apply?**

- **Suggestion:** Start with a small, invite-only adult audience that enjoys interest communities and participatory governance, using English for the first version. Confirm target locations and audience policy before a public launch. This is a scope proposal, not a conclusion about regulatory requirements.

**Q27 — What should the product feel like, and which devices matter first?**

- **Suggestion:** Build a responsive web app with a shared potato identity, common navigation, and each society's name, image, and governance-type badge. Society discovery should distinguish identity, interests, membership conditions, and type. Use plain wording for governing powers, competition rules, and moderation outcomes. Prioritize mobile layouts, labels, keyboard access, and readable contrast. Earlier login aesthetic suggestions remain provisional.

**Q28 — What backend, hosting, budget, and operational setup should we use?**

- **Suggestion:** Keep the existing SvelteKit/Svelte/Tailwind frontend and use managed authentication plus a relational database when connecting real data. Store society identity separately from governance type and enforce membership/office permissions per society on the server. Add election and competition infrastructure when their slices are built. Choose services after deployment region, budget, provider preferences, and operator access are known; use isolated development data before live operation.

**Q29 — How will we decide whether the pilot works?**

- **Suggestion:** Evaluate the feature currently being piloted. For the entry flow, check whether people understand society versus type and can create or join successfully. Later assess conversation, return visits, understanding of governing powers, completion of votes, moderation outcomes, and completed challenges between societies. Collect short qualitative feedback alongside aggregate usage. Choose numerical targets per checkpoint rather than requiring unavailable features to produce metrics.

### Shared world and competition

**Q30 — Where do societies meet, and can members join conversations in societies they do not belong to?**

- **Suggestion:** Use one home, search, profile system, and eventually a challenge hub. Let people discover societies across governance types and join public conversations where the host's participation rules permit it. Keep governing rights tied to membership in the particular society. Display the host space's rules and apply them equally to visitors. Shared event discussion uses common platform rules with neutral moderation.

**Q31 — Which societies can challenge each other, and who can commit each society?**

- **Suggestion:** Allow any two distinct societies to challenge one another, including Kingdom-versus-Kingdom and Republic-versus-Republic. Let members propose ideas, with official acceptance by an authorized leader or delegate of each participating society according to its type. The same person should not approve both sides of a contest. Record identical accepted terms and let commitments survive leadership turnover under the event rules. Start with two-society events; multi-society tournaments can follow later.

**Q32 — Which competition format should the first competition checkpoint support?**

- **Suggestion:** Start with one small, scheduled online team competition using an existing game or activity, with registration, shared discussion, manually submitted evidence, and reviewed results inside Potatoverse. Choose the actual activity with the first participants before building format-specific screens. Keep the core challenge flow flexible enough for future esports, creative contests, quizzes, or real-life events. Do not assume every format uses teams, identical scoring, or a game integration.

**Q33 — Who decides results and resolves competition disputes when the societies disagree?**

- **Suggestion:** Agree on scoring, evidence, a neutral referee, and a dispute deadline before accepting the challenge. Show submitted results as provisional until confirmed. Use a different reviewer for an appeal and define the final authority before the event starts. Keep a readable result history and make corrections visible. Moderate event conversation separately from adjudicating the competition itself.

**Q34 — How are participants selected, and what prevents population imbalance or switching sides during an event?**

- **Suggestion:** For the first team format, use equal team sizes, open sign-ups, a published selection method, and a roster deadline. A person represents only one society in an event, even if a member of several participating societies. Snapshot representation at roster lock; leaving the represented society ends participation rather than transferring the person to its opponent mid-event. Changing the society viewed in the app does not change representation. Publish substitution/withdrawal rules and welcome spectators from any society.

**Q35 — What do winners receive, and how much event organization happens inside the app?**

- **Suggestion:** Start with event-specific recognition, profile/team badges, and shared result history, without money, paid entry, political power, or a universal score across incompatible activities. Let organizers coordinate the chosen online activity externally while Potatoverse handles challenges, participation, and results. If offline events are selected later, first define hosts, locations, attendance/privacy, cancellations, and participant support. Choose any broader society standings only after agreeing on comparable scoring and anti-farming rules.

### Society creation and incremental delivery

**Q36 — What creation limits and joining options should the first society flow support?**

- **Suggestion:** Let verified members create a society directly, with one unfinished founding society per account during the pilot to keep abandoned setups manageable. Start with public society profiles and open joining; private societies and membership approval can come later. Save interrupted creation as a draft and distinguish previewing a draft from opening a governed society. Define how creating a society affects the founder's membership under Q01.

**Q37 — Can an established society change governance type, and who can archive it?**

- **Suggestion:** Allow type changes while still a draft. Lock the type once members join for the first working version. Later conversion should require an agreed member decision and an explicit plan for offices, active elections, rules, and commitments. Allow a founder to discard an unused draft, but require a published society-level process to archive an established society. Founding an active society must not give someone an unexplained permanent deletion right.

**Q38 — What identifies a society, and how should country or interest affect discovery?**

- **Suggestion:** Require a name, unique address, short description, and governance type. Keep country, language, interests, and image optional. Allow multiple societies for the same country or topic; never assign governance type from a person's location. Display enough information to distinguish similar names, and do not label country-themed societies as official national representatives without a separate policy.

**Q39 — Which governance type should we implement first, and how much customization should founders get?**

- **Suggestion:** Implement Kingdom first because scoped leader-to-moderator appointments provide a smaller first authority model than the proposed representative advancement ladder. Its initial ruler-selection and founding rules still need to be decided. Make multiple Kingdom societies work, then add Republic using the same account, society, membership, and content foundations. Offer predefined type behavior with limited approved settings; defer custom governing systems and type conversion. Republic can come first instead if its election experience is the feature we most want to validate.

## Development breakdown — incremental, with a proposed order

The current repository contains the SvelteKit starter page, Svelte 5, Vite, and Tailwind CSS 4. No custom interface or authentication implementation was observed during this planning pass. These are future tasks. Incremental delivery is confirmed; the order below and the first governance type are suggestions pending Q10/Q39.

Complete and review one small user outcome at a time. Each row can be split further. The whole table is a direction for development, not a single release commitment. A local interface prototype can proceed with labeled assumptions; a working feature must resolve the decisions needed to persist data and enforce its actual permissions.

| Order / task | Deliverable | Main decisions needed | Completion evidence |
| --- | --- | --- | --- |
| T00 — Choose the next small outcome | Confirm the first checkpoint and note only the assumptions it needs. | Q01, Q10, Q36, Q39 | A short implementation brief identifies the next reviewable flow and its boundaries. |
| T01 — Entry-flow prototype | Login/sign-up interface, society directory/details, create/join paths, and a minimal society home in one responsive app shell. Use clearly labeled example data until real services are connected. | Q10, Q27, Q38; visible assumptions from T00 | A person can distinguish society from type and walk through creation/joining; empty, invalid, and interrupted states are reviewable on mobile and by keyboard. No simulated sign-in is presented as real authentication. |
| T02 — Persistent accounts and societies | Connect sign-in/recovery, profiles, society drafts, and scoped membership. A society can enter an explicitly labeled founding state only once its limited founding controls are implemented; ordinary governing powers follow in T04. | Q01, Q07, Q11–Q12, Q25–Q26, Q28, Q36–Q39 | Creation survives reload, repeated submissions do not duplicate societies, and joining references the correct society. Two societies of the same type have separate members and permissions. Types without implemented founding controls remain draft-only. |
| T03 — Conversation with basic moderation | Add communities, text posts/comments/reactions, discovery, rules, scoped moderation, reports/appeals, and basic notifications in small successive tasks. Use only explicitly approved founding authority until ordinary leaders take office. | Q02–Q04, Q07–Q08, Q13–Q17, Q25, Q30 | Members can converse and obtain a moderation outcome; a founder or moderator in one society cannot act in another. The agreed rules and reviewer route are available before real member posting. |
| T04 — First governance type | Complete the selected type's leadership selection, offices, appointments, terms, and handover. Kingdom is the current recommendation, not a confirmed choice. | Q04–Q05, Q07, Q09, Q18–Q24, Q39 for Kingdom; Q06 if Republic is selected first | One society completes its approved founding-to-leadership cycle and another of the same type does so independently. Expired or revoked powers no longer work. |
| T05 — First shared challenge | Build proposal/acceptance first, then entries, shared discussion, reviewed results, and disputes for one selected format. | Q04, Q12, Q16–Q17, Q25, Q30–Q35 | Two named societies accept the same terms and complete an event. Same-type competitors work if Q31 approves them; no second governance type is required for that pairing. |
| T06 — Next governance type | Add Republic if Kingdom was first, reusing accounts, society identity, membership, and content while implementing its distinct election ladder and powers. | Q03–Q09, Q15, Q18–Q24, Q39 | New societies can select the completed type; existing societies retain their behavior. Each society's elections and offices remain independent. |
| T07 — Further expansion | Select the next needed feature from feedback: more governance types, richer content, new competition formats, or other approved improvements. | Questions relevant to that selected feature | One additional outcome is implemented and verified without requiring the entire long-term vision at once. |

Before opening each working slice to real users, verify its recovery, permissions, moderation coverage, and relevant operator actions. Pilot readiness belongs to each increment, rather than waiting until every type and competition format is finished.

Reuse election code only where the types actually share rules. Do not hardcode exactly two societies, assign leadership by governance type alone, or build a general-purpose governance engine before the concrete supported models need it.

Before coding each task, write a small issue with the user outcome, linked confirmed decisions, entry/exit states, necessary permissions, affected screens/data, and acceptance criteria. Ask only unresolved questions that materially affect that task. Advance when that task is verified; expand or reorder later tasks as the owner decides.

## Cross-flow scenarios — verify when the relevant feature is ready

- A new person understands society versus governance type, creates or joins a named society, and returns later without repeating onboarding.
- A founder saves a society draft, resumes it, selects a supported type, invites members, and follows the first-leadership policy without silently receiving permanent governing powers.
- Two societies serving the same country can choose different types; two societies of the same type retain distinct leaders, elections, members, and moderation powers.
- A request using another society's identifier cannot reuse a founder's, king's, president's, or moderator's authority from the current society.
- Duplicate society creation, an unavailable type, an abandoned draft, and founder departure each produce a clear, recoverable outcome.
- When social features are available, a member joins a community, posts, receives a reply, and returns to the conversation.
- An interrupted invite/sign-in flow resumes at the correct community and explains any society membership conflict.
- A casual member can participate without voting or running for office.
- Members of different societies can discover each other, interact in shared spaces, and follow the same challenge without changing accounts or affiliation.
- An official challenge is proposed, accepted by both sides, filled fairly, completed, and recorded; changed terms require renewed agreement.
- If approved under Q31, two societies of the same governance type can compete. Adding another supported governance type does not change their existing memberships or results.
- A contested score, absent referee, no-show, or cancelled event follows the agreed process without either society rewriting the result unilaterally.
- A competitor switches affiliation or a ruler leaves office during an accepted event; roster and commitment rules preserve a predictable outcome.
- A Kingdom member becomes king through the agreed process, appoints a consenting moderator, and later hands over authority correctly.
- A Republic member wins a community seat, wins the next approved office, and leaves a community vacancy that is filled predictably.
- Someone loses an election and retains ordinary membership rights; duplicate ballot attempts and eligibility changes are handled consistently.
- A moderator is reported, the case goes to an appropriate reviewer, an affected member appeals, and the outcome updates access correctly.
- A leader resigns, expires, or is suspended while an action is open; their next request cannot use stale authority.
- A community has no leader, a contest has no candidates or insufficient turnout, or an outage spans voting close; the published fallback is followed.
- A member switches societies, leaves, recovers an account, or requests deletion without unexpectedly duplicating votes, abandoning offices, or exposing private records.
- Each critical journey works on mobile and by keyboard, including errors, empty states, confirmation steps, and recovery.

## Next planning checkpoint

Focus on Q01, Q10, Q36, and Q39 to choose the next small deliverable and its assumptions. The suggested starting point is the entry-flow prototype in T01, with society creation/discovery built around separate society identities and governance types. Resolve other questions as their features approach implementation. Record answers, remove their suggestions, update affected flows, and proceed one part at a time.
