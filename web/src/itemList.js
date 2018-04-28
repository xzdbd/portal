import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import Item from './item'

const styles = {
  root: {
    flexGrow: 1,
    paddingTop: 40,
    minHeight: 500
  },
  dividerGrid: {
    paddingTop: 0,
    borderTop: '4px solid #4179F7'
  },
  container: {
    width: '100%'
  }
}

function ItemList (props) {
  const { classes } = props
  return (
    <div className={classes.root}>
      <Grid container spacing={40} justify='center' className={classes.container}>
        <Grid item xs={12} className={classes.dividerGrid} />
        <Grid item xs={12}>
          <Grid container spacing={40} justify='center'>
            <Grid item xs={8}>
              <Item />
            </Grid>
            <Grid item xs={8}>
              <Item />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  )
}

ItemList.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(ItemList)
