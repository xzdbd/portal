import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import HeaderBar from './header'
import Search from './search'
import ItemList from './itemList'
import Footer from './footer'

const styles = theme => ({
  root: {
    flexGrow: 1
  },
  paper: {
    padding: theme.spacing.unit * 2,
    textAlign: 'center',
    color: theme.palette.text.secondary
  }
})

function FullWidthGrid (props) {
  const { classes } = props

  return (
    <div className={classes.root}>
      <Grid container spacing={0} justify='center'>
        <Grid item xs={12} >
          <HeaderBar />
        </Grid>
        <Grid item xs={12}>
          <Search />
        </Grid>
        <Grid item xs={12}>
          <ItemList />
        </Grid>
        <Grid item xs={12}>
          <Footer />
        </Grid>
      </Grid>
    </div>
  )
}

FullWidthGrid.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(FullWidthGrid)
