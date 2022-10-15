it('displays one h1 title by default', () => {
    cy.get('h1').should('have.length', 1)

    cy.get('h1').first().should('have.text', 'version')
    cy.get('p').last().should('have.text', 'Build User: giuliohome')
  })
